package repository

import (
	"context"
	"encoding/json"
	"presence-service/internal"
	"presence-service/internal/config"
	"presence-service/internal/domain"
	"slices"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	logger, _ = internal.WireLogger()
)

type RedisClient struct {
	rdb *redis.Client // Thread safe.
}

func New(cfg *config.Config) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
        Addr:     cfg.Redis.HostAddress,
        Password: "", // no password set
        DB:       cfg.Redis.DB, 
    })
	return &RedisClient{rdb: rdb}
}

func (r *RedisClient) Close() {
	if err := r.rdb.Close(); err != nil {
		logger.Error(
			"unable to close redis",
			zap.String("trace", err.Error()),
		)
	}
}

func (r *RedisClient) UpdateUser(ctx context.Context, arg domain.HeartbeatRequest) error {
	user, err := r.GetUser(ctx, arg.UserID)
	if err != nil {
		// New user.
		user = domain.User{
			UserID: arg.UserID,
			Servers: make([]string, 0),
		}
	}

	// New device.
	if !slices.Contains(user.Servers, arg.Server) {
		user.Servers = append(user.Servers, arg.Server)
	}

	user.LastHeartbeat = time.Now().Format(time.RFC3339)

	p, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := r.rdb.Set(ctx, user.UserID, p, 30 * time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) GetUser(ctx context.Context, userID string) (domain.User, error) {
	val, err := r.rdb.Get(ctx, userID).Result()
	if err != nil {
		return domain.User{}, err
	}

	p := new(domain.User)
	if err := json.Unmarshal([]byte(val), p); err != nil {
		return domain.User{}, err
	}
	return *p, err
}