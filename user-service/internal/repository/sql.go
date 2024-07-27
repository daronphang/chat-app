package repository

import (
	"context"
	"user-service/internal/domain"
)

func (q *Querier) CreateUser(ctx context.Context, arg domain.NewUser) (domain.UserMetadata, error) {
	stmt := `
	INSERT INTO user_metadata (
	user_id, email, display_name
	) VALUES (
	 $1, $2, $3
	)
	RETURNING user_id, email, display_name, created_at
	`
	row := q.db.QueryRow(ctx, stmt, arg.UserID, arg.Email, arg.DisplayName)

	var i domain.UserMetadata
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.DisplayName,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Querier) UpdateUser(ctx context.Context, arg domain.UserMetadata) error {
	stmt := `
	UPDATE user_metadata SET 
	email = $1,
	display_name = $2
	WHERE user_id = $3
	`

	_, err := q.db.Exec(ctx, stmt, arg.UserID, arg.Email, arg.DisplayName)
	if err != nil {
		return err
	}

	return nil
}