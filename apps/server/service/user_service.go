package service

import (
	"context"
	"log"
	"saldo-server/domain"
	"saldo-server/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u *UserService) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := u.UserRepository.GetByID(ctx, userID)
	if err != nil {
		log.Println("Failed to get user by id")
		return nil, err
	}
	return user, nil
}
