package tasks_postgres_repository

import (
	"time"
	"todo-list/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserId int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserId,
	)
}

func taskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	taskDomains := make([]domain.Task, len(taskModels))

	for i, taskModel := range taskModels {
		taskDomains[i] = domain.Task{
			ID:           taskModel.ID,
			Version:      taskModel.Version,
			Title:        taskModel.Title,
			Description:  taskModel.Description,
			Completed:    taskModel.Completed,
			CreatedAt:    taskModel.CreatedAt,
			CompletedAt:  taskModel.CompletedAt,
			AuthorUserID: taskModel.AuthorUserId,
		}
	}

	return taskDomains
}
