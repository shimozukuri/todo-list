package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shimozukuri/todo-list/internal/core/domain"
	core_logger "github.com/shimozukuri/todo-list/internal/core/logger"
	core_http_request "github.com/shimozukuri/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/shimozukuri/todo-list/internal/core/transport/http/response"
	core_http_types "github.com/shimozukuri/todo-list/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" swaggertype:"string" example:"Petr Petrovich"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string" example:"+71112223344"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("'FullName' can't be NULL")
		}

		fullNameLength := len([]rune(*r.FullName.Value))
		if fullNameLength < 3 || fullNameLength > 100 {
			return fmt.Errorf("'FullName' length must be between 3 and 100")
		}
	}

	if r.PhoneNumber.Set && r.PhoneNumber.Value != nil {
		phoneNumberLength := len([]rune(*r.PhoneNumber.Value))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("'PhoneNumber' length must be between 10 and 15")
		}

		if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
			return fmt.Errorf("'PhoneNumber' must starts with '+'")
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser 	godoc
// @Summary 	Update User
// @Description Update user by ID
// @Description ### Field update behavior (three-state logic):
// @Description 1. **Field omitted**: `phone_number` is ignored; the existing database value remains unchanged.
// @Description 2. **Value provided**: `"phone_number": "+71112223344"` updates the database value.
// @Description 3. **Explicit null**: `"phone_number": null` sets the database value to `NULL`.
// @Description **Restriction**: `full_name` cannot be `null`.
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "user ID"
// @Param		request body PatchUserRequest true "PatchUser request body"
// @Success 	200 {object} PatchUserResponse "Success to update user"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 	409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	var request PatchUserRequest
	if err = core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
