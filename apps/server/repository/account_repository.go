package repository

import (
	"context"
	"log"
	"saldo-server/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	Pool *pgxpool.Pool
}

func NewAccountRepository(pool *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		Pool: pool,
	}
}

func (a *AccountRepository) Create(ctx context.Context, userID string, ca *domain.CreateAccount) (*domain.Account, error) {
	query := "INSERT INTO accounts (name, user_id) VALUES ($1, $2) RETURNING *"

	var account domain.Account
	err := a.Pool.QueryRow(ctx, query, ca.Name, userID).Scan(
		&account.ID,
		&account.Name,
		&account.UserID,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *AccountRepository) ListByUserID(ctx context.Context, userID string) (*[]domain.Account, error) {
	query := "SELECT * FROM accounts WHERE user_id = $1 ORDER BY created_at DESC"
	rows, err := a.Pool.Query(ctx, query, userID)
	if err != nil {
		log.Printf("Failed to query accounts %v\n", err)
		return nil, err
	}
	defer rows.Close()

	accounts := []domain.Account{}
	for rows.Next() {
		account := domain.Account{}
		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.UserID,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			log.Printf("Failed to retreive accounts %v\n", err)
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return &accounts, nil
}
