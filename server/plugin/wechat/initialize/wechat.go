package initialize

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/common"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/pkg/tools"
	"go.uber.org/zap"
	"time"
)

// InitWechatConfig 微信配置初始化
func InitWechatConfig() {
	var wechatConfig model.Wechat
	rdb := global.GlobalConfig.Rdb
	ctx := context.Background()
	_, err := tools.SetRedisStrResult[model.Wechat](rdb, ctx, common.GetWechatConfigKey(), wechatConfig, -time.Second)
	if err != nil {
		global.GlobalConfig.Log.Error("InitWechatConfig error", zap.Error(err))
		return
	}
}
