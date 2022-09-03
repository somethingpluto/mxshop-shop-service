package initialize

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"go.uber.org/zap"
	"inventory_service/global"
)

func InitRedis() {
	redisConfig := global.ServiceConfig.Redis
	client := goredislib.NewClient(&goredislib.Options{Addr: fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port), Password: redisConfig.Password})
	pool := goredis.NewPool(client)
	global.Redsync = redsync.New(pool)
	zap.S().Infof("redsync初始化成功 \n")
}
