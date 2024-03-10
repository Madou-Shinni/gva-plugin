package common

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/global"
)

const (
	WechatConfigKey      = "we_chat_config:%s"
	WeChatAccessTokenKey = "we_chat_access_token:%s"
	WeChatJsApiTicketKey = "we_chat_js_api_ticket:%s"
)

// GetWeChatAccessTokenKey 获取微信AccessToken的key
func GetWeChatAccessTokenKey(appid string) string {
	return fmt.Sprintf(WeChatAccessTokenKey, appid)
}

// GetWeChatJsApiTicketKey 获取微信JsApiTicket的key
func GetWeChatJsApiTicketKey(appid string) string {
	return fmt.Sprintf(WeChatJsApiTicketKey, appid)
}

// GetWechatConfigKey 获取微信配置key
func GetWechatConfigKey() string {
	return fmt.Sprintf(WechatConfigKey, global.GlobalConfig.ID)
}
