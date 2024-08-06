package repository

import (
	"context"
	"message-service/internal/domain"
	"time"
)

func (q *Querier) GetLatestMessages(ctx context.Context, channelID string) ([]domain.Message, error) {
	stmt := `
	SELECT messageId, channelId, senderId, messageType, content, createdAt 
	FROM message WHERE channelId = ?
	ORDER BY messageId DESC
	LIMIT 100
	`
	scanner := q.session.Query(
		stmt,
		channelID,
 	).WithContext(ctx).Iter().Scanner()

	var items []domain.Message
	for scanner.Next() {
		var i domain.Message
		var createdAt int64
		if err := scanner.Scan(
			&i.MessageID,
			&i.ChannelID,
			&i.SenderID,
			&i.MessageType,
			&i.Content,
			&createdAt,
		); err != nil {
			return nil, err
		}
		createdAt /= 1000
		i.CreatedAt = time.Unix(createdAt, 0).Format("2006-01-02T15:04:05Z07:00")
		items = append(items, i)
	}
	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *Querier) GetPreviousMessages(ctx context.Context, arg domain.PrevMessageRequest) ([]domain.Message, error) {
	stmt := `
	SELECT messageId, channelId, senderId, messageType, content, createdAt 
	FROM message WHERE channelId = ? AND messageId < ?
	ORDER BY messageId DESC
	LIMIT 100
	`
	scanner := q.session.Query(
		stmt,
		arg.ChannelID,
		arg.LastMessageID,
 	).WithContext(ctx).Iter().Scanner()

	var items []domain.Message
	for scanner.Next() {
		var i domain.Message
		var createdAt int64
		if err := scanner.Scan(
			&i.MessageID,
			&i.ChannelID,
			&i.SenderID,
			&i.MessageType,
			&i.Content,
			&createdAt,
		); err != nil {
			return nil, err
		}
		createdAt /= 1000
		i.CreatedAt = time.Unix(createdAt, 0).Format("2006-01-02T15:04:05Z07:00")
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
	INSERT INTO message (messageId, channelId, senderId, messageType, content, createdAt) 
	VALUES 
	(?, ?, ?, ?, ?, ?)
	`

	createdAt, err := time.Parse("2006-01-02T15:04:05Z07:00", arg.CreatedAt)
	if err != nil {
		return err
	}

	if err := q.session.Query(
		stmt,
		arg.MessageID,
		arg.ChannelID,
		arg.SenderID,
		arg.MessageType,
		arg.Content,
		createdAt,
	).WithContext(ctx).Exec(); err != nil {
		return err
	}
	return nil
}

// func (q *Querier) GetUserRelations(ctx context.Context, userID string) ([]string, error) {
// 	stmt := `SELECT relationId FROM user_relation WHERE userId = ?`
// 	scanner := q.session.Query(
// 		stmt,
// 		userID,
//  	).WithContext(ctx).Iter().Scanner()

// 	 var items []string
// 	 for scanner.Next() {
// 		 var relationID string
// 		 if err := scanner.Scan(&relationID); err != nil {
// 			 return nil, err
// 		 }
// 		 items = append(items, relationID)
// 	 }
// 	 // scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
// 	 if err := scanner.Err(); err != nil {
// 		 return nil, err
// 	 }
// 	 return items, nil
// }

// func (q *Querier) GetUserIdsAssociatedToChannel(ctx context.Context, channelID string) ([]string, error) {
// 	stmt := `SELECT userId FROM channel WHERE channelId = ?`
// 	scanner := q.session.Query(
// 		stmt,
// 		channelID,
//  	).WithContext(ctx).Iter().Scanner()

// 	var items []string
// 	for scanner.Next() {
// 		var userID string
// 		if err := scanner.Scan(&userID); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, userID)
// 	}
// 	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// func (q *Querier) GetChannelsAssociatedToUserID(ctx context.Context, userID string) ([]string, error) {
// 	stmt := `SELECT channelId FROM user WHERE userId = ?`
// 	scanner := q.session.Query(
// 		stmt,
// 		userID,
//  	).WithContext(ctx).Iter().Scanner()

// 	var items []string
// 	for scanner.Next() {
// 		var channelID string
// 		if err := scanner.Scan(&channelID); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, channelID)
// 	}
// 	// scanner.Err() closes the iterator, so scanner nor iter should be used afterwards.
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }



// // when adding users to channel, need to add to both tables channel and user.
// //
// // Channel table is used for broadcasting events to users in a group.
// //
// // User table is used to determine which channels belong to a user.
// func (q *Querier) AddUserIDsToChannel(ctx context.Context, channelID string, userIDs []string) error {
// 	stmt1 := `
// 	INSERT INTO channel(channelId, userId, createdAt) 
// 	VALUES 
// 	(?, ?, toTimestamp(now()))
// 	`
// 	stmt2 := `
// 	INSERT INTO user(userId, channelId, createdAt) 
// 	VALUES 
// 	(?, ?, toTimestamp(now()))
// 	`
// 	stmt3 := `
// 	INSERT INTO user_relation(userId, relationId, createdAt) 
// 	VALUES 
// 	(?, ?, toTimestamp(now()))
// 	`


// 	b := q.session.NewBatch(gocql.LoggedBatch).WithContext(ctx)

// 	for _, userID := range userIDs {
// 		b.Entries = append(b.Entries, gocql.BatchEntry{
// 			Stmt:       stmt1,
// 			Args:       []interface{}{channelID, userID},
// 			Idempotent: true,
// 		})
// 	}

// 	for _, userID := range userIDs {
// 		b.Entries = append(b.Entries, gocql.BatchEntry{
// 			Stmt:       stmt2,
// 			Args:       []interface{}{userID, channelID},
// 			Idempotent: true,
// 		})
// 	}

// 	if len(userIDs) == 2 {
// 		b.Entries = append(b.Entries, gocql.BatchEntry{
// 			Stmt:       stmt3,
// 			Args:       []interface{}{userIDs[0], userIDs[1]},
// 			Idempotent: true,
// 		})
// 		b.Entries = append(b.Entries, gocql.BatchEntry{
// 			Stmt:       stmt3,
// 			Args:       []interface{}{userIDs[1], userIDs[0]},
// 			Idempotent: true,
// 		})

// 	}

// 	if err := q.session.ExecuteBatch(b); err != nil {
// 		return err
// 	}
// 	return nil
// }