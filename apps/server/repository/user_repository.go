package repository

import (
	"context"
	"fmt"
	"log"
	"saldo-server/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		Pool: pool,
	}
}

func (u *UserRepository) GetByID(ctx context.Context, userID string) (*domain.User, error) {
	query := "SELECT * FROM users where id = $1"
	rows, err := u.Pool.Query(ctx, query, userID)
	if err != nil {
		log.Printf("Error query user %v\n", err)
		return nil, err
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
	if err != nil {
		log.Printf("Error parse to row user %v\n", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := "SELECT * FROM users where phone = $1"
	rows, err := u.Pool.Query(ctx, query, phone)
	if err != nil {
		log.Printf("Error query get user by phone %v\n", err)
		return nil, err
	}
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])

	if err != nil {
		log.Printf("Error parse to row user %v\n", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) HasAccountID(ctx context.Context, userID string, accountID string) error {
	query := "SELECT COUNT(*) FROM accounts a where a.id = $1 AND a.user_id = $2"
	row := u.Pool.QueryRow(ctx, query, accountID, userID)

	var count int
	err := row.Scan(&count)

	if err != nil {
		log.Printf("Error parsing count user accounts %v\n", err)
		return err
	}

	if count == 0 {
		log.Printf("Error parsing count user accounts %v\n", err)
		return fmt.Errorf("account %s is not associated with user %s", accountID, userID)
	}

	return nil
}
