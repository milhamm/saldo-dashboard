package domain

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	UserID    string `json:"user_id,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

type CreateAccount struct {
	Name string `json:"name"`
}
