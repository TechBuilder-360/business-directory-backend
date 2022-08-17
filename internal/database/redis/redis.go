package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	defaultExpirationTime = (time.Hour * 24) * 30 // 30 days
)

// Client used to make requests to redis
type Client struct {
	*redis.Client
	ttl       time.Duration
	namespace string
}

var redisClient *Client

// NewClient is a client constructor.
func NewClient(connectionURL, password, namespace string) *Client {
	log.Info("connecting to redis client")

	c := redis.NewClient(&redis.Options{
		Addr:        connectionURL,
		Password:    password, // no password set
		DB:          0,
		DialTimeout: 15 * time.Second,
		MaxRetries:  10, // use default DB
	})

	// Test redis connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if _, err := c.Ping(ctx).Result(); err != nil {
		log.Panic("unable to connect to redis: %s", err)
	}

	log.Info("connected to redis client")
	client := &Client{
		Client:    c,
		ttl:       defaultExpirationTime,
		namespace: namespace,
	}

	setRedisClient(client)

	return client
}

func setRedisClient(client *Client) {
	redisClient = client
}

func RedisClient() *Client {
	return redisClient
}

func (c *Client) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := c.Client.Ping(ctx).Result()
	return err
}

func (c *Client) Set(key string, value interface{}, duration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Set(ctx, key, value, duration).Err()
}

func (c *Client) Get(key string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Get(ctx, key).Result()
}

func (c *Client) Delete(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Del(ctx, key).Result()
}

func (c *Client) Exists(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	i, err := c.Client.Exists(ctx, key).Result()
	return i >= 1, err
}
