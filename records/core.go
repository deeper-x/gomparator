package records

import (
	"log"

	"github.com/go-redis/redis/v7"
)

// Repository struct
type Repository struct {
	Client *redis.Client
}

// NewRepository create client instance
func NewRepository() Repository {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return Repository{Client: client}
}

// SetKey is a wrapper for the client set
func (c Repository) SetKey(key string, value string) {
	err := c.Client.Set(key, value, 0).Err()

	if err != nil {
		log.Println(err)
	}
}

// GetKey is a wrapper for the client get
func (c Repository) GetKey(key string) string {
	val, err := c.Client.Get(key).Result()

	if err != nil {
		log.Println(err)
	}

	return val
}

// RPush is a redis rpush list wrapper
func (c Repository) RPush(value string) {
	c.Client.RPush("alist", value)
}

// GetAll returns all list items
func (c Repository) GetAll() *redis.StringSliceCmd {
	return c.Client.LRange("alist", 0, -1)
}
