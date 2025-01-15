package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"saldo-server/domain"
	"saldo-server/repository"
	"strings"
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
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, domain.NewGenericError(http.StatusNotFound, "User not found", fmt.Errorf("user %s not found", userID))
		}
		return nil, err
	}
	return user, nil
}
