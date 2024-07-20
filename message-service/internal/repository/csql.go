package repository

import (
	"context"
	"message-service/internal/domain"

	"github.com/gocql/gocql"
)

func (q *Querier) GetMessages(ctx context.Context, channelID string) ([]domain.Message, error) {
	stmt := `SELECT * FROM message WHERE channelId = ?`
	scanner := q.session.Query(
		stmt,
		channelID,
 	).WithContext(ctx).Iter().Scanner()

	var items []domain.Message
	for scanner.Next() {
		var i domain.Message
		if err := scanner.Scan(
			&i.MessageID,
			&i.PreviousMessageID,
			&i.ChannelID,
			&i.SenderID,
			&i.Type,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *Querier) CreateMessage(ctx context.Context, arg domain.Message) error {
	stmt := `
	INSERT INTO message (messageId, previousMessageId, channelId, senderId, type, content, createdAt) 
	VALUES 
	(?, ?, ?, ?, ?, ?, ?)
	`
	if err := q.session.Query(
		stmt,
		arg.MessageID,
		arg.PreviousMessageID,
		arg.ChannelID,
		arg.SenderID,
		arg.Type,
		arg.Content,
		arg.CreatedAt,
	).WithContext(ctx).Exec(); err != nil {
		return err
	}
	return nil
}

// when adding users to channel, need to add to both tables channel and user.
// Channel table is used for broadcasting events to users in a group.
// User table is used to determine which channels belong to a user.
// For newly created channels, need to check if its 1-on-1 chat or group chat (len(users) > 2).
// This is to ensure 1-on-1 chats have a single channelId, in the event both users try to create
// a new chat at the same time (edge case).
func (q *Querier) AddUsersToChannel(ctx context.Context, channelID string, users []string) error {
	stmt1 := `
	INSERT INTO channel(channelId, userId, createdAt) 
	VALUES 
	(?, ?, ?)
	`
	// For 1-on-1 chats.
	stmt2 := `
	INSERT INTO user(userId, user2Id, channelId, createdAt) 
	VALUES 
	(?, ?, ?, ?)
	`
	// For group chats.
	stmt3 := `
	INSERT INTO user(userId, channelId, createdAt) 
	VALUES 
	(?, ?, ?)
	`

	b := q.session.NewBatch(gocql.LoggedBatch).WithContext(ctx)

	b.Entries = append(b.Entries, gocql.BatchEntry{
		Stmt:       "INSERT INTO example.batches (pk, ck, description) VALUES (?, ?, ?)",
		Args:       []interface{}{1, 2, "1.2"},
		Idempotent: true,
	})

	if err := q.session.ExecuteBatch(b); err != nil {
		return err
	}
	return nil
	// dateof(now())
}