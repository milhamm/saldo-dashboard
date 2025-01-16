package service

import (
	"context"
	"saldo-server/domain"
	"saldo-server/repository"
)

type AccountService struct {
	AccountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
	}
}

func (a *AccountService) ListAccounts(ctx context.Context, userID string) (*[]domain.Account, error) {
	accounts, err := a.AccountRepository.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (a *AccountService) CreateAccount(ctx context.Context, userID string, ca *domain.CreateAccount) (*domain.Account, error) {
	account, err := a.AccountRepository.Create(ctx, userID, ca)
	if err != nil {
		return nil, err
	}
	return account, nil
}
