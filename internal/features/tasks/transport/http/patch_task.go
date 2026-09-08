package tasks_transport_http

import (
	"fmt"
	"net/http"
	"todo-list/internal/core/domain"
	core_logger "todo-list/internal/core/logger"
	core_http_request "todo-list/internal/core/transport/http/request"
	core_http_response "todo-list/internal/core/transport/http/response"
	core_http_types "todo-list/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
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
