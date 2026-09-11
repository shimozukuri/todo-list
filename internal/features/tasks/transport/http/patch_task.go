package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/shimozukuri/todo-list/internal/core/domain"
	core_logger "github.com/shimozukuri/todo-list/internal/core/logger"
	core_http_request "github.com/shimozukuri/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/shimozukuri/todo-list/internal/core/transport/http/response"
	core_http_types "github.com/shimozukuri/todo-list/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title" swaggertype:"string" example:"Buy chicken"`
	Description core_http_types.Nullable[string] `json:"description" swaggertype:"string" example:"1 kilogram"`
	Completed   core_http_types.Nullable[bool]   `json:"completed" swaggertype:"boolean" example:"true"`
}

func (r *PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("'Title' can't be NULL")
		}

		titleLength := len([]rune(*r.Title.Value))
		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf("'Title' length must be between 1 and 100")
		}
	}

	if r.Description.Set && r.Description.Value != nil {
		descriptionLength := len([]rune(*r.Description.Value))
		if descriptionLength < 1 || descriptionLength > 1000 {
			return fmt.Errorf("'Description' length must be between 1 and 1000")
		}
	}

	if r.Completed.Set && r.Completed.Value == nil {
		return fmt.Errorf("'Completed' can't be NULL")
	}

	return nil
}

type PatchTaskResponse TasksDTOResponse

// PatchTask	godoc
// @Summary 	Update Task
// @Description Update task by ID
// @Description ### Field update behavior (three-state logic):
// @Description 1. **Field omitted**: `description` is ignored; the existing database value remains unchanged.
// @Description 2. **Value provided**: `"description": "2 liters"` updates the database value.
// @Description 3. **Explicit null**: `"description": null` sets the database value to `NULL`.
// @Description **Restriction**: `title` and `completed` can't be `null`.
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "task ID"
// @Param		request body PatchTaskRequest true "PatchTask request body"
// @Success 	200 {object} PatchTaskResponse "Success to update task"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 	409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get taskID path value",
		)

		return
	}

	var request PatchTaskRequest

	if err = core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	taskPatch := taskPatchFromRequest(request)

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch task",
		)

		return
	}

	response := PatchTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}
