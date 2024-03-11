package methods

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"onlineStoreBackend/constants"
)

func GetDatabase() (db *gorm.DB, err error) {
	dsn := "host=" + constants.PostgresData.Host +
		" port=" + constants.PostgresData.Port[rand.Intn(2)] +
		" user=" + constants.PostgresData.User +
		" dbname=" + constants.PostgresData.DBName +
		" password=" + constants.PostgresData.Password +
		" sslmode=" + constants.PostgresData.SSLMode

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	return db, err
}

func GetRedis(ctx context.Context) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "sanya",
		DB:       0,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		return client, err
	}

	return client, err
}
