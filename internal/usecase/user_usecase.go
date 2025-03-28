package usecase

import (
	"errors"
	"go-college/internal/domain/entity"
	"go-college/internal/domain/repository"

	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(name string) (*entity.User, error)
	GetUserByID(userID string) (*entity.User, error)
	UpdateUser(userID string, name string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

// CreateUser ユーザー作成
func (u *userUsecase) CreateUser(name string) (*entity.User, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	authToken, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:        userID.String(),
		AuthToken: authToken.String(),
		Name:      name,
		HighScore: 0,
		Coin:      0,
	}

	err = u.userRepo.InsertUser(user)
	return user, err
}

// GetUserByID ユーザー情報取得
func (u *userUsecase) GetUserByID(userID string) (*entity.User, error) {
	user, err := u.userRepo.SelectUserByPrimaryKey(userID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser ユーザー情報更新
func (u *userUsecase) UpdateUser(userID string, name string) error {
	user, err := u.userRepo.SelectUserByPrimaryKey(userID)
	if err != nil || user == nil {
		return errors.New("user not found")
	}
	user.Name = name
	return u.userRepo.UpdateUserByPrimaryKey(user)
}
