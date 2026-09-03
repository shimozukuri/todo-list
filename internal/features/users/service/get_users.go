package users_service

import (
	"context"
	"fmt"
	"todo-list/internal/core/domain"
	core_errors "todo-list/internal/core/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
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

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return users, nil
}
