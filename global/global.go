package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	// Logger *logrus.Logger
	Db      *gorm.DB
	RedisDB *redis.Client
)
