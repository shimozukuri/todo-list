package tasks_transport_http

import (
	"time"
	"todo-list/internal/core/domain"
)

type TasksDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
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
