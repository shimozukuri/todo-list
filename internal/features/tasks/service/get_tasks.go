package tasks_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
	core_errors "todo-list/internal/core/errors"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID, limit, offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be a positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be a positive: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	taskDomains, err := s.taskRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	return taskDomains, nil
}
