package tasks_transport_http

import (
	"time"

	"github.com/shimozukuri/todo-list/internal/core/domain"
)

type TasksDTOResponse struct {
	ID           int        `json:"id" example:"1"`
	Version      int        `json:"version" example:"1"`
	Title        string     `json:"title" example:"Buy CocaCola"`
	Description  *string    `json:"description" example:"2 liters"`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"null"`
	CompletedAt  *time.Time `json:"completed_at" example:"2006-01-01T00:00:00.000000Z"`
	AuthorUserID int        `json:"author_user_id" example:"1"`
}

func taskDTOFromDomain(task domain.Task) TasksDTOResponse {
	return TasksDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOFromDomains(tasks []domain.Task) []TasksDTOResponse {
	tasksDTO := make([]TasksDTOResponse, len(tasks))

	for i, task := range tasks {
		tasksDTO[i] = taskDTOFromDomain(task)
	}

	return tasksDTO
}
