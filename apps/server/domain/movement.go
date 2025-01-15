package domain

type MovementType string

const (
	MovementTypeWithdraw MovementType = "withdraw"
	MovementTypeTransfer MovementType = "transfer"
	MovementTypeTopup    MovementType = "top_up"
	MovementTypePayment  MovementType = "payment"
	MovementTypeOthers   MovementType = "others"
)

type Movement struct {
	ID           string       `json:"id"`
	Amount       int64        `json:"amount"`
	Fee          int64        `json:"fee"`
	MovementType MovementType `json:"movement_type"`
	AccountID    string       `json:"account_id,omitempty"`
	CreatedAt    *int64       `json:"created_at,omitempty"`
	UpdatedAt    *int64       `json:"updated_at,omitempty"`
}

type CreateMovement struct {
	Amount       int64        `json:"amount"`
	Fee          int64        `json:"fee"`
	MovementType MovementType `json:"movement_type"`
}

/*
   - Timestamp
   - Tanggal
   - Nama (Receipent)
   - Movement type
   - IN/OUT
   - Nominal
   - Admin
*/
