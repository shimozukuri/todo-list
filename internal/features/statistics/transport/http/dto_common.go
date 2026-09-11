package statistics_transport_http

import (
	"github.com/shimozukuri/todo-list/internal/core/domain"
)

type StatisticsDTOResponse struct {
	TasksCreated               int      `json:"tasks_created" example:"20"`
	TasksCompleted             int      `json:"tasks_completed" example:"2"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate" example:"10"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time" example:"5m30s"`
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
