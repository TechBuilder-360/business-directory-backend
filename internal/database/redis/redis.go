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

func (c *Client) Get(key string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	result, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	return &result, err
}

func (c *Client) Exists(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	res, err := c.Client.Exists(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}

	return res > 0, err
}

func (c *Client) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	key = fmt.Sprintf("%s-%s", c.namespace, key)
	return c.Client.Del(ctx, key).Err()
}
