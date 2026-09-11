package users_postgres_repository

import "github.com/shimozukuri/todo-list/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainFromModel(model UserModel) domain.User {
	return domain.NewUser(
		model.ID,
		model.Version,
		model.FullName,
		model.PhoneNumber,
	)
}

func userDomainsFromModels(models []UserModel) []domain.User {
	userDomains := make([]domain.User, len(models))

	for i, model := range models {
		userDomains[i] = userDomainFromModel(model)
	}

	return userDomains
}
