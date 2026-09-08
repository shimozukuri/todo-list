package statistics_service

import (
	"context"
	"fmt"
	"time"
	"todo-list/internal/core/domain"
	core_errors "todo-list/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'from' must be before 'to': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statisticsRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get tasks from repository: %w", err)
	}

	statisticDomain := calcStatistics(tasks)

	return statisticDomain, nil
}

func calcStatistics(tasks []domain.Task) domain.Statistics {
	if len(tasks) == 0 {
		return domain.NewStatistics(
			0,
			0,
			nil,
			nil,
		)
	}

	tasksCreated := len(tasks)

	tasksCompleted := 0
	var totalCompletionDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			tasksCompleted++
		}

		completionDuration := task.CompletionDuration()
		if completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}

	tasksCompletedRate := new(float64(tasksCompleted) / float64(tasksCreated) * 100)

	var tasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletionDuration != 0 {
		tasksAverageCompletionTime = new(totalCompletionDuration / time.Duration(tasksCompleted))
	}

	return domain.NewStatistics(
		tasksCreated,
		tasksCompleted,
		tasksCompletedRate,
		tasksAverageCompletionTime,
	)
}
