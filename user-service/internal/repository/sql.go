package repository

import (
	"context"
	"time"
	"user-service/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	var createdAt pgtype.Timestamp
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.DisplayName,
		&createdAt,
	)
	i.CreatedAt = createdAt.Time.String()
	return i, err
}

func (q *Querier) GetUser(ctx context.Context, arg string) (domain.UserMetadata, error) {
	stmt := `
	SELECT user_id, email, display_name, created_at
	FROM user_metadata
	WHERE email = $1
	`

	row := q.db.QueryRow(ctx, stmt, arg)

	var i domain.UserMetadata
	var createdAt pgtype.Timestamp
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.DisplayName,
		&createdAt,
	)
	i.CreatedAt = createdAt.Time.String()
	return i, err
}

func (q *Querier) UpdateUser(ctx context.Context, arg domain.UserMetadata) error {
	stmt := `
	UPDATE user_metadata SET 
	email = $1,
	display_name = $2
	WHERE user_id = $3
	`

	_, err := q.db.Exec(ctx, stmt, arg.Email, arg.DisplayName, arg.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) AddFriend(ctx context.Context, arg domain.NewFriend) error {
	stmt := `
	INSERT INTO friendship (
	user_id, friend_id, display_name
	) VALUES (
	 $1, $2, $3
	)
	`
	_, err := q.db.Exec(ctx, stmt, arg.UserID, arg.FriendID, arg.DisplayName)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) GetFriends(ctx context.Context, arg string) ([]domain.Friend, error) {
	stmt := `
	SELECT 
	FS.friend_id AS friend_id,
	UM.email AS email,
	FS.display_name AS display_name
	FROM friendship AS FS
	INNER JOIN user_metadata AS UM ON FS.friend_id = UM.user_id
	WHERE FS.user_id = $1
	`

	rows, err := q.db.Query(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []domain.Friend
	for rows.Next() {
		var i domain.Friend
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
			&i.DisplayName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, err
}

func (q *Querier) CreateUserToChannelAssociation(ctx context.Context, arg domain.NewChannel) error {
	var rows [][]interface{}
	createdAt, _ := time.Parse("2006-01-02T15:04:05Z07:00", arg.CreatedAt)
	for _, userID := range arg.UserIDs {
		rows = append(rows, []interface{}{userID, arg.ChannelID, createdAt})
	}

	_, err := q.db.CopyFrom(
		ctx,
		pgx.Identifier{"user_to_channel"},
		[]string{"user_id", "channel_id", "created_at"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}
	return nil
}

func (q *Querier) CreateGroupChannel(ctx context.Context, arg domain.NewChannel) error {
	stmt := `
	INSERT INTO group_channel (
	channel_id, group_name
	) VALUES (
	 $1, $2, $3
	)
	`
	createdAt, _ := time.Parse("2006-01-02T15:04:05Z07:00", arg.CreatedAt)
	_, err := q.db.Exec(ctx, stmt, arg.ChannelID, arg.ChannelName, createdAt)
	if err != nil {
		return err
	}
	return nil
}

func (q *Querier) GetUsersAssociatedToChannel(ctx context.Context, arg string) ([]domain.UserContact, error) {
	stmt := `
	SELECT 
	UTC.user_id AS user_id,
	UM.email AS email
	FROM user_to_channel AS UTC
	INNER JOIN user_metadata AS UM ON UM.user_id = UTC.user_id
	WHERE UTC.channel_id = $1
	`

	rows, err := q.db.Query(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []domain.UserContact
	for rows.Next() {
		var i domain.UserContact
		if err := rows.Scan(&i.UserID, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, err
}

func (q *Querier) GetChannelsAssociatedToUser(ctx context.Context, arg string) ([]domain.Channel, error) {
	stmt := `
	SELECT 
	UTC.channel_id AS channel_id,
	CASE WHEN GC.group_name IS NOT NULL THEN GC.group_NAME WHEN FS.display_name IS NOT NULL THEN FS.display_name ELSE UM.email END AS channel_name,
	UTC.created_at as created_at	
	FROM
	user_to_channel AS UTC
	LEFT JOIN group_channel AS GC ON GC.channel_id = UTC.channel_id
	LEFT JOIN friendship AS FS ON FS.user_id = UTC.user_id AND POSITION(FS.friend_id IN UTC.channel_id) > 0
	LEFT JOIN user_metadata AS UM ON UM.user_id != UTC.user_id AND POSITION(UM.user_id IN UTC.channel_id) > 0
	WHERE UTC.user_id = $1
	`

	rows, err := q.db.Query(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []domain.Channel
	var ts pgtype.Timestamp
	for rows.Next() {
		var i domain.Channel
		if err := rows.Scan(
			&i.ChannelID,
			&i.ChannelName,
			&ts,
		); err != nil {
			return nil, err
		}
		i.CreatedAt = ts.Time.String()
		items = append(items, i)
	}
	return items, err
}

// To fetch both user contacts and unknown users of channels the user is in.
func (q *Querier) GetUsersAssociatedToTargetUser(ctx context.Context, arg string) ([]string, error) {
	// TODO:
	stmt := `
	SELECT user_id FROM user_to_channel
	WHERE channel_id = $1
	`

	rows, err := q.db.Query(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []string
	for rows.Next() {
		var i string
		if err := rows.Scan(&i); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, err
}