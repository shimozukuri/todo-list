package tasks_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf(
			"validate task domain: %w",
			err,
		)
	}

	taskDomain, err := s.taskRepository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"create task: %w",
			err,
		)
	}

	return taskDomain, nil
}
