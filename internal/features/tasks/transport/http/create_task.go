package tasks_transport_http

import (
	"net/http"

	"github.com/shimozukuri/todo-list/internal/core/domain"
	core_logger "github.com/shimozukuri/todo-list/internal/core/logger"
	core_http_request "github.com/shimozukuri/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/shimozukuri/todo-list/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100" example:"Buy CocaCola"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000" example:"2 liters"`
	AuthorUserID int     `json:"author_user_id" validate:"required" example:"1"`
}

type CreateTaskResponse TasksDTOResponse

// CreateTask 	godoc
// @Summary 	Create Task
// @Description Create new task
// @Tags 		tasks
// @Accept 		json
// @Produce 	json
// @Param 		request body CreateTaskRequest true "CreateTask request body"
// @Success 	201 {object} CreateTaskResponse "Success to create new task"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404 {object} core_http_response.ErrorResponse "Author not found"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks [post]
func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var req CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	taskDomain := domainFromDTO(req)

	taskDomain, err := h.tasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create task",
		)

		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(
		dto.Title,
		dto.Description,
		dto.AuthorUserID,
	)
}
