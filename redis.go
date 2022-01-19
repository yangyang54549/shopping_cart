package shopping_cart

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
)

type RedisDb struct {
	Db  *redis.Client
	Key string
}

var ctx = context.Background()

func (t *RedisDb) New(key interface{}, Db interface{}) Storage {
	db := new(RedisDb)
	db.Db = Db.(*redis.Client)
	db.Key = key.(string)
	return db
}

func (t *RedisDb) CleanOne(rowIds string) error {
	if len(rowIds) == 0 {
		return t.Db.Del(ctx, t.Key).Err()
	}
	return t.Db.HDel(ctx, t.Key, rowIds).Err()
}

func (t *RedisDb) Clean() error {
	return t.Db.Del(ctx, t.Key).Err()
}

func (t *RedisDb) Save(rowId string, data *ShoppingCart) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return t.Db.HSet(ctx, t.Key, rowId, b).Err()
}

func (t *RedisDb) Get(rowId string) (*ShoppingCart, error) {
	str, err := t.Db.HGet(ctx, t.Key, rowId).Result()
	if err != nil {
		return nil, err
	}

	var data = new(ShoppingCart)
	if err := json.Unmarshal([]byte(str), data); err != nil {
		return nil, err
	}

	return data, nil
}

func (t *RedisDb) IsRow(rowId string) (bool, error) {
	_, err := t.Db.HGet(ctx, t.Key, rowId).Result()

	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *RedisDb) GetAll() (res []ShoppingCart, err error) {
	var data ShoppingCart

	rows, err := t.Db.HGetAll(ctx, t.Key).Result()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		err := json.Unmarshal([]byte(row), &data)
		if err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}
