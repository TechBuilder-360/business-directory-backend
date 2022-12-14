package middlewares

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	log "github.com/sirupsen/logrus"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/redis"
	"time"
)

var CacheClient *cache.Client

func ResponseCache() {
	ringOpt := &redis.RingOptions{
		Addrs: map[string]string{
			"server": configs.Instance.RedisURL,
		},
		DB:       configs.Instance.RedisDB,
		Password: configs.Instance.RedisPassword,
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(redis.NewAdapter(ringOpt)),
		cache.ClientWithTTL(10*time.Second),
		cache.ClientWithRefreshKey(configs.Instance.RedisCacheRefresh),
	)
	if err != nil {
		log.Fatalf("failed to setup response cache. %s", err.Error())
	}

	CacheClient = cacheClient
}
