package common

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/pkg/tools"
)

// GetWechatAccessToken 从中控服务器获取微信AccessToken
func GetWechatAccessToken() string {
	ctx := context.Background()
	rdb := global.GlobalConfig.Rdb
	appId := global.GlobalConfig.Wechat.AppID

	// 从中控服务器获取access token
	accessToken, err := tools.GetRedisStrResult[string](rdb, ctx, GetWeChatAccessTokenKey(appId))
	if err != nil {
		return ""
	}
	return accessToken
}

// GetWechatPublicAccessToken 从中控服务器获取微信公众号AccessToken
func GetWechatPublicAccessToken() string {
	ctx := context.Background()
	rdb := global.GlobalConfig.Rdb
	appId := global.GlobalConfig.Wechat.PubWxConfig.AppID

	// 从中控服务器获取access token
	accessToken, err := tools.GetRedisStrResult[string](rdb, ctx, GetWeChatAccessTokenKey(appId))
	if err != nil {
		return ""
	}
	return accessToken
}

// GetWechatPublicJsApiTicket 从中控服务器获取微信公众号JsApiTicket
func GetWechatPublicJsApiTicket() string {
	ctx := context.Background()
	rdb := global.GlobalConfig.Rdb
	appId := global.GlobalConfig.Wechat.PubWxConfig.AppID

	// 从中控服务器获取JsApiTicket
	jsApiTicket, err := tools.GetRedisStrResult[string](rdb, ctx, GetWeChatJsApiTicketKey(appId))
	if err != nil {
		return ""
	}
	return jsApiTicket
}
