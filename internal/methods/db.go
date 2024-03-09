package methods

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"onlineStoreBackend/constants"
)

func GetDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres",
		"host="+constants.PostgresData.Host+
			" port="+constants.PostgresData.Port[rand.Intn(2)]+
			" user="+constants.PostgresData.User+
			" dbname="+constants.PostgresData.DBName+
			" password="+constants.PostgresData.Password)
	return db, err
}

func GetRedis(ctx context.Context) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "sanya",
		DB:       0,
	})

	result, err := client.Ping(ctx).Result()
	if err != nil {
		return client, err
	}

	log.Print(result)

	return client, err
}
