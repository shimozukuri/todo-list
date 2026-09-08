package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"
	core_logger "todo-list/internal/core/logger"
	core_http_request "todo-list/internal/core/transport/http/request"
	core_http_response "todo-list/internal/core/transport/http/response"
)

type GetStatisticsResponse StatisticsDTOResponse

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'userID'/'from'/'to' query param",
		)

		return
	}

	statisticDomain, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics: %w",
		)

		return
	}

	response := GetStatisticsResponse(statisticDTOFromDomain(statisticDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		UserIDQueryParamKey = "user_id"
		FromQueryParamKey   = "from"
		ToQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, UserIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'userID' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, FromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, ToQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}
