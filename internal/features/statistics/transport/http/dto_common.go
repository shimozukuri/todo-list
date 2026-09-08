package statistics_transport_http

import (
	"todo-list/internal/core/domain"
)

type StatisticsDTOResponse struct {
	TasksCreated               int      `json:"tasks_created"`
	TasksCompleted             int      `json:"tasks_completed"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func statisticDTOFromDomain(statistic domain.Statistics) StatisticsDTOResponse {
	var avgTime *string
	if statistic.TasksAverageCompletionTime != nil {
		avgTime = new(statistic.TasksAverageCompletionTime.String())
	}

	return StatisticsDTOResponse{
		TasksCreated:               statistic.TasksCreated,
		TasksCompleted:             statistic.TasksCompleted,
		TasksCompletedRate:         statistic.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}
