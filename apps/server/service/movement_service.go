package service

import (
	"context"
	"log"
	"net/http"
	"saldo-server/domain"
	"saldo-server/repository"
)

type MovementService struct {
	MovementRepository *repository.MovementRepository
	UserRepository     *repository.UserRepository
}

func NewMovementService(movementRepository *repository.MovementRepository, userRepository *repository.UserRepository) *MovementService {
	return &MovementService{
		MovementRepository: movementRepository,
		UserRepository:     userRepository,
	}
}

func (m *MovementService) CreateMovement(ctx context.Context, userID string, accountID string, payload *domain.CreateMovement) (*domain.Movement, error) {
	err := m.UserRepository.HasAccountID(ctx, userID, accountID)
	if err != nil {
		return nil, domain.NewGenericError(http.StatusNotFound, "Account not found", err)
	}

	movement, err := m.MovementRepository.Create(ctx, accountID, payload.Amount, payload.Fee, payload.MovementType)
	if err != nil {
		log.Println("Failed to create movement")
		return nil, err
	}

	return movement, nil
}

func (m *MovementService) ListMovements(ctx context.Context, userID string) ([]domain.Movement, error) {
	movements, err := m.MovementRepository.ListByUserID(ctx, userID)
	if err != nil {
		log.Println("Failed to retreive list movements")
		return nil, err
	}
	return movements, nil
}
