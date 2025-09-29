package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nicitapa/firstProgect/internal/configs"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"time"
)

type Cache struct {
	rdb    *redis.Client
	logger zerolog.Logger
}

func NewCache(client *redis.Client, logger zerolog.Logger) *Cache {
	return &Cache{
		rdb:    client,
		logger: logger,
	}
}
func (c *Cache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	rawU, err := json.Marshal(value)
	if err != nil {
		fmt.Println("error during marshal:", err)
		return err
	}

	if err = c.rdb.Set(ctx, c.formatKey(key), rawU, duration).Err(); err != nil {
		fmt.Println("error during set:", err)
		return err
	}

	return nil
}

func (c *Cache) Get(ctx context.Context, key string, response interface{}) error {
	val, err := c.rdb.Get(ctx, c.formatKey(key)).Result()
	if err != nil {
		fmt.Println("error during get:", err)
		return err
	}

	if err = json.Unmarshal([]byte(val), response); err != nil {
		fmt.Println("error during unmarshal:", err)
		return err
	}

	return nil
}

func (c *Cache) formatKey(key string) string {
	return fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerName, key)
}
