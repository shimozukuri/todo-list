package tasks_service

import (
	"context"
	"fmt"
	"github.com/shimozukuri/todo-list/internal/core/domain"
)

func (s *TasksService) PatchTask(
	ctx context.Context,
	taskID int,
	task domain.TaskPatch,
) (domain.Task, error) {
	domainTask, err := s.taskRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}

	if err = domainTask.ApplyPatch(task); err != nil {
		return domain.Task{}, fmt.Errorf("apply task patch: %w", err)
	}

	patchedTask, err := s.taskRepository.PatchTask(ctx, taskID, domainTask)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task: %w", err)
	}

	return patchedTask, nil
}
