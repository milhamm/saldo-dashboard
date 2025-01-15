package repository

import (
	"context"
	"log"
	"saldo-server/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovementRepository struct {
	Pool *pgxpool.Pool
}

func NewMovementRepository(pool *pgxpool.Pool) *MovementRepository {
	return &MovementRepository{
		Pool: pool,
	}
}

func (m *MovementRepository) Create(ctx context.Context, accountID string, amount int64, fee int64, movementType domain.MovementType) (*domain.Movement, error) {
	query := "INSERT INTO movements (amount, fee, type, account_id) VALUES ($1, $2, $3, $4) RETURNING *"

	var movement domain.Movement
	err := m.Pool.QueryRow(ctx, query, amount, fee, movementType, accountID).Scan(
		&movement.ID,
		&movement.Amount,
		&movement.Fee,
		&movement.MovementType,
		&movement.AccountID,
		&movement.CreatedAt,
		&movement.UpdatedAt,
	)

	if err != nil {
		log.Println("Failed to insert movement")
		return nil, err
	}

	return &movement, nil
}

func (m *MovementRepository) ListByUserID(ctx context.Context, userID string) ([]domain.Movement, error) {
	query := `
        SELECT
            m.id,
            m.amount,
            m.fee,
            m."type",
            m.created_at,
            m.updated_at
        FROM
            movements m
        LEFT JOIN accounts a ON a.id = m.account_id
        WHERE a.user_id = $1
    `
	rows, err := m.Pool.Query(ctx, query, userID)
	if err != nil {
		log.Println("Failed to query movement by user id")
		return nil, err
	}
	defer rows.Close()

	movements := []domain.Movement{}

	for rows.Next() {
		movement := domain.Movement{}
		err := rows.Scan(
			&movement.ID,
			&movement.Amount,
			&movement.Fee,
			&movement.MovementType,
			&movement.CreatedAt,
			&movement.UpdatedAt,
		)
		if err != nil {
			log.Println("Failed to collect rows movement by id")
			return nil, err
		}

		movements = append(movements, movement)
	}

	return movements, nil
}
