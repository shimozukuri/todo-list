package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"todo-list/internal/core/domain"
	core_errors "todo-list/internal/core/errors"
	core_postgres_pool "todo-list/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	taskID int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, title, description, 
			completed, created_at, completed_at, author_user_id
		FROM todolist.tasks
		WHERE id = $1;
	`

	row := r.pool.QueryRow(ctx, query, taskID)

	var taskModel TaskModel

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorUserId,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d': %w",
				taskID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("scan error %w", err)
	}

	taskDomain := taskDomainFromModel(taskModel)
	
	return taskDomain, nil
}
