package repository

import "go-college/internal/domain/entity"

type UserRepository interface {
	InsertUser(user *entity.User) error
	SelectUserByAuthToken(authToken string) (*entity.User, error)
	SelectUserByPrimaryKey(userID string) (*entity.User, error)
	UpdateUserByPrimaryKey(user *entity.User) error
}
