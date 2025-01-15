package service

import (
	"context"
	"errors"
	"net/http"
	"os"
	"saldo-server/common/password"
	"saldo-server/domain"
	"saldo-server/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUnauthorized = errors.New("invalid phone or password")
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (a *AuthService) Login(ctx context.Context, u *domain.LoginRequest) (string, error) {
	user, err := a.UserRepository.GetByPhone(ctx, u.Phone)
	if err != nil {
		return "", domain.NewGenericError(http.StatusUnauthorized, "Unauthorized", ErrUnauthorized)
	}

	match, err := password.ComparePasswordAndHash(u.Password, user.Password)
	if err != nil {
		return "", err
	}

	if !match {
		return "", domain.NewGenericError(http.StatusUnauthorized, "Unauthorized", ErrUnauthorized)
	}

	claims := &domain.JWTClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
