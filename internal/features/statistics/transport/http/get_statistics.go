package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/shimozukuri/todo-list/internal/core/logger"
	core_http_request "github.com/shimozukuri/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/shimozukuri/todo-list/internal/core/transport/http/response"
)

type GetStatisticsResponse StatisticsDTOResponse

// GetStatistics godoc
// @Summary 	 Get statistics
// @Description  Get statistics with optional filtering by author user ID and date range
// @Tags         statistics
// @Produce      json
// @Param        user_id query int false "Filter by author user ID"
// @Param        from query string false "Start date for statistics calculation (inclusive), format: YYYY-MM-DD"
// @Param        to query string false "End date for statistics calculation (exclusive), format: YYYY-MM-DD"
// @Success 	 200 {object} GetStatisticsResponse "Success to get statistics"
// @Failure 	 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		 /statistics [get]
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
