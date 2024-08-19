package repository

import (
	"context"
	"time"
	"user-service/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
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

	// Postgres returns timestamp in local timezone format.
	i.CreatedAt = createdAt.Time.Format(time.RFC3339)
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
	i.CreatedAt = createdAt.Time.Format(time.RFC3339)
	return i, err
}

func (q *Querier) GetUserById(ctx context.Context, arg string) (domain.UserMetadata, error) {
	stmt := `
	SELECT user_id, email, display_name, created_at
	FROM user_metadata
	WHERE user_id = $1
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
	i.CreatedAt = createdAt.Time.Format(time.RFC3339)
	return i, err
}

func (q *Querier) GetUsers(ctx context.Context, arg []string) ([]domain.UserMetadata, error) {
	stmt := `
	SELECT user_id, email, display_name, created_at
	FROM user_metadata
	WHERE user_id = ANY($1::varchar[])
	`

	rows, err := q.db.Query(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []domain.UserMetadata
	for rows.Next() {
		var i domain.UserMetadata
		var createdAt pgtype.Timestamp
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
			&i.DisplayName,
			&createdAt,
		); err != nil {
			return nil, err
		}
		i.CreatedAt = createdAt.Time.Format(time.RFC3339)
		items = append(items, i)
	}
	return items, err
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
	_, err := q.db.Exec(ctx, stmt, arg.UserID, arg.FriendID, arg.FriendName)
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
	UM.display_name as display_name,
	FS.display_name AS friend_name
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
			&i.FriendName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, err
}

func (q *Querier) CreateUserToChannelAssociation(ctx context.Context, arg domain.Channel) error {
	var rows [][]interface{}
	createdAt, _ := time.Parse(time.RFC3339, arg.CreatedAt)
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

func (q *Querier) UpdateLastReadMessage(ctx context.Context, arg domain.LastReadMessage) error {
	stmt := `
	UPDATE user_to_channel SET 
	last_message_id = $1
	WHERE user_id = $2
	AND channel_id = $3
	`

	_, err := q.db.Exec(ctx, stmt, arg.LastMessageID, arg.UserID, arg.ChannelID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) CreateGroupChannel(ctx context.Context, arg domain.Channel) error {
	stmt := `
	INSERT INTO group_channel (
	channel_id, group_name, created_at
	) VALUES (
	 $1, $2, $3
	)
	`
	createdAt, _ := time.Parse(time.RFC3339, arg.CreatedAt)
	_, err := q.db.Exec(ctx, stmt, arg.ChannelID, arg.ChannelName, createdAt)
	if err != nil {
		return err
	}
	return nil
}

func (q *Querier) GetGroupChannel(ctx context.Context, arg string) (domain.Channel, error) {
	stmt := `
	SELECT channel_id, group_name, created_at
	FROM 
	group_channel 
	WHERE channel_id = $1
	`

	row := q.db.QueryRow(ctx, stmt, arg)

	var i domain.Channel
	var createdAt pgtype.Timestamp
	err := row.Scan(
		&i.ChannelID,
		&i.ChannelName,
		&createdAt,
	)
	i.CreatedAt = createdAt.Time.Format(time.RFC3339)
	return i, err
}

func (q *Querier) RemoveGroupMembers(ctx context.Context, arg domain.GroupMembers) error {
	stmt := `
	DELETE FROM user_to_channel
	WHERE channel_id = $1 AND user_id = ANY($2::varchar[]);
	`

	_, err := q.db.Exec(ctx, stmt, arg.ChannelID, pq.Array(arg.UserIDs))
	if err != nil {
		return err
	}
	return nil
}

func (q *Querier) RemoveGroup(ctx context.Context, arg string) error {
	stmt1 := `
	DELETE FROM user_to_channel
	WHERE channel_id = $1;
	`
	stmt2 := `
	DELETE FROM group_channel 
	WHERE channel_id = $1;
	`
	db := q.db.(*pgxpool.Pool)
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = db.Exec(ctx, stmt1, arg)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, stmt2, arg)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
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
	UTC.user_id as user_id,
	UTC.channel_id AS channel_id,
	CASE WHEN GC.group_name IS NOT NULL THEN GC.group_NAME ELSE '' END AS channel_name,
	UTC.created_at as created_at,
	COALESCE(UTC.last_message_id, 0) AS last_message_id	
	FROM
	user_to_channel AS UTC
	LEFT JOIN group_channel AS GC ON GC.channel_id = UTC.channel_id
	WHERE UTC.channel_id IN (SELECT channel_id FROM user_to_channel WHERE user_id = $1)
	ORDER BY channel_id
	`

	rows, err := q.db.Query(ctx, stmt, arg)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []domain.Channel
	for rows.Next() {
		var i domain.Channel
		var ts pgtype.Timestamp
		var userID string
		if err := rows.Scan(
			&userID,
			&i.ChannelID,
			&i.ChannelName,
			&ts,
			&i.LastMessageID,
		); err != nil {
			return nil, err
		}
		i.CreatedAt = ts.Time.Format(time.RFC3339)

		if len(items) > 0 && items[len(items)-1].ChannelID == i.ChannelID {
			prev := &items[len(items)-1]
			prev.UserIDs = append(prev.UserIDs, userID)

			if userID == arg {
				prev.LastMessageID = i.LastMessageID
			}
		} else {
			i.UserIDs = append(i.UserIDs, userID)
			items = append(items, i)
		}
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

func (q *Querier) GetUsersContactsMetadata(ctx context.Context, arg []string) ([]domain.UserContact, error) {
	stmt := `
	SELECT user_id, email
	FROM user_metadata
	WHERE user_id = ANY($1::varchar[]);
	`

	rows, err := q.db.Query(ctx, stmt, pq.Array(arg))
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []domain.UserContact
	for rows.Next() {
		var i domain.UserContact
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, err
}