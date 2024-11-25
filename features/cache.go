package features

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	"RouteHub.Service.Dashboard/features/configuration"
	"github.com/redis/go-redis/v9"
)

var (
	onceConfigureCache sync.Once
	cacheService       *CacheService
)

type CacheService struct {
	logger      *slog.Logger
	cacheConfig *configuration.CacheDatabaseConfig
	options     *redis.Options
}

const (
	errCacheServiceNotInitialized = "cache service not initialized"
	errRedisCheck                 = "error checking redis"
	errRedisPing                  = "error pinging redis"
	errConfigureRedis             = "error configuring redis client"
	errLoadingCacheConfig         = "error loading cache config"
	errGettingRedisClient         = "error getting redis client"
	errRedisClientIsNil           = "redis client is nil"
	errRedisAddress               = "redis address is empty"
)

var (
	ErrCacheServiceNotInitialized = errors.New(errCacheServiceNotInitialized)
	ErrRedisCheck                 = errors.New(errRedisCheck)
	ErrRedisPing                  = errors.New(errRedisPing)
	ErrConfigureRedis             = errors.New(errConfigureRedis)
	ErrLoadingCacheConfig         = errors.New(errLoadingCacheConfig)
	ErrGettingRedisClient         = errors.New(errGettingRedisClient)
	ErrRedisClientIsNil           = errors.New(errRedisClientIsNil)
	ErrRedisAddress               = errors.New(errRedisAddress)
)

func NewCacheService(cdc *configuration.CacheDatabaseConfig, logger *slog.Logger) (*CacheService, error) {
	ctx := context.Background()
	if cdc == nil {
		err := ErrCacheServiceNotInitialized
		logger.ErrorContext(ctx, errLoadingCacheConfig, "error", err)
		return nil, err
	}

	var initErr error
	onceConfigureCache.Do(func() {
		cacheService = &CacheService{
			logger:      logger,
			cacheConfig: cdc,
		}

		if err := cacheService.configureOptions(); err != nil {
			initErr = ErrConfigureRedis
			logger.ErrorContext(ctx, errConfigureRedis, "error", err)
			return
		}

		if err := cacheService.checkRedis(); err != nil {
			initErr = ErrRedisCheck
			logger.ErrorContext(ctx, errRedisCheck, "error", err)
			return
		}

		if cdc.ClearOnStart {
			if err := flushCache(logger); err != nil {
				initErr = err
				return
			}
		}
	})

	if initErr != nil {
		return nil, initErr
	}

	return cacheService, nil
}

func (cs *CacheService) checkRedis() error {
	ctx := context.Background()

	client, err := GetCacheClient()
	if err != nil {
		return ErrGettingRedisClient
	}

	if client == nil {
		return ErrRedisClientIsNil
	}

	defer client.Close()

	res, err := client.Ping(ctx).Result()
	if err != nil && res != "PONG" {
		return ErrRedisPing
	}

	return nil
}

func (cs *CacheService) configureOptions() error {
	addr := cs.cacheConfig.GetAddress()
	cs.logger.Debug("Redis Config", "address", addr)

	if addr == "" {
		return ErrRedisAddress
	}

	cs.options = &redis.Options{
		Addr: addr,
	}

	if cs.cacheConfig.Password != "" {
		cs.options.Password = cs.cacheConfig.Password
	}

	if cs.cacheConfig.DB != 0 {
		cs.options.DB = cs.cacheConfig.DB
	}

	return nil
}

func GetCacheService() (*CacheService, error) {
	if cacheService == nil {
		return nil, ErrCacheServiceNotInitialized
	}

	return cacheService, nil
}

func GetCacheClient() (*redis.Client, error) {
	if cacheService == nil {
		return nil, ErrRedisClientIsNil
	}

	return redis.NewClient(cacheService.options), nil
}

func flushCache(logger *slog.Logger) error {
	client, err := GetCacheClient()
	if err != nil {
		logger.ErrorContext(context.Background(), errRedisPing, "error", ErrRedisPing, "err", err)
		return err
	}
	defer client.Close()

	client.FlushAll(context.Background())
	return nil
}
