package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type CartsRepository struct{}

func (s CartsRepository) PostProducts(ctx context.Context, client *redis.Client, userID int, ids []interface{}) (err error) {
	err = client.RPush(ctx, strconv.Itoa(userID), ids...).Err()
	return err
}

func (s CartsRepository) DeleteCart(ctx context.Context, client *redis.Client, userID int) (err error) {
	err = client.Del(ctx, strconv.Itoa(userID)).Err()
	return err
}

func (s CartsRepository) GetProductsFromCart(ctx context.Context, client *redis.Client, userID int) (ids []string, err error) {
	ids, err = client.LRange(ctx, strconv.Itoa(userID), 0, -1).Result()
	return ids, err
}
