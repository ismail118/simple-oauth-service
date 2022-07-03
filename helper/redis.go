package helper

import (
	"context"
	"github.com/go-redis/redis/v8"
	"simple-oauth-service/constanta"
	"simple-oauth-service/errors"
	"time"
)

func Set(ctx context.Context, rdb *redis.Client, key string, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func SetWithExp(ctx context.Context, rdb *redis.Client, key string, value string, minute int) error {
	ttl := time.Duration(minute) * time.Minute

	err := rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(ctx context.Context, rdb *redis.Client, key string) (string, error) {
	result, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", errors.NewValidationErrors(constanta.AuthorizeCodeNotExist)
	}
	if err != nil {
		return "", err
	}

	return result, nil
}

func Delete(ctx context.Context, rdb *redis.Client, key string) {
	rdb.Del(ctx, key)
}
