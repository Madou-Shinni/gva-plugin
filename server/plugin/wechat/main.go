package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/jobs"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/pkg/tools"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/router"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type wechatPlugin struct{}

func CreateWechatPlug(AppId, AppSecret, PublicAppId, PublicAppSecret string, rdb *redis.Client, log *zap.Logger) *wechatPlugin {
	global.GlobalConfig.Wechat = tools.NewWxConfig(AppId, AppSecret, PublicAppId, PublicAppSecret)
	global.GlobalConfig.Rdb = rdb
	global.GlobalConfig.Log = log
	return &wechatPlugin{}
}

func (*wechatPlugin) Register(group *gin.RouterGroup) {
	jobs.CronInit()
	router.RouterGroupApp.InitWechatRouter(group)
}

func (*wechatPlugin) RouterPath() string {
	return "wechat"
}
