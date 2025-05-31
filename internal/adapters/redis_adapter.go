package adapters

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/frostnzx/go-myapi-assignment/internal/core"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type RedisProfileRepository struct {
	client *redis.Client
}

func NewRedisProfileRepository(client *redis.Client) core.ProfileRepository {
	return &RedisProfileRepository{client: client}
}

func (r *RedisProfileRepository) Save(profile core.Profile) error {
	// generate a unique key for the profile using UUID
	id := uuid.New().String()
	key := "profile:" + id
	data, err := json.Marshal(profile)
	if err != nil {
		return errors.New("failed to marshal profile data")
	}

	// save the profile data in Redis with the generated key
	ctx := context.Background() // Use a context for Redis operations
	if err := r.client.Set(ctx, key, data, 0).Err(); err != nil {
		return errors.New("failed to save profile to Redis")
	}
	return nil
}

func (r *RedisProfileRepository) GetAll() ([]core.Profile, error) {
	ctx := context.Background()
	var profiles []core.Profile

	iter := r.client.Scan(ctx, 0, "profile:*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()

		data, err := r.client.Get(ctx, key).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			return nil, errors.New("failed to get profile data from Redis")
		}

		var profile core.Profile
		if err := json.Unmarshal([]byte(data), &profile); err != nil {
			return nil, errors.New("failed to deserialize profile data")
		}

		profiles = append(profiles, profile)
	}

	if err := iter.Err(); err != nil {
		return nil, errors.New("failed to iterate over Redis keys")
	}

	return profiles, nil
}
