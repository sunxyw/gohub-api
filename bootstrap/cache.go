// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"gohub/pkg/cache"
	"gohub/pkg/config"
)

// SetupCache 缓存
func SetupCache() {

	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.Get[string]("redis.host"), config.Get[string]("redis.port")),
		config.Get[string]("redis.username"),
		config.Get[string]("redis.password"),
		config.Get[int]("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
