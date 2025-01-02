package repository

import (
	"context"
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
