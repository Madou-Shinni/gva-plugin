package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/api"
	"github.com/gin-gonic/gin"
)

type WechatRouter struct{}

func (s *WechatRouter) InitWechatRouter(Router *gin.RouterGroup) {
	wechatRouter := Router
	wechatApi := api.ApiGroupApp.WechatApi
	{
		wechatRouter.GET("jsapi", wechatApi.GetJsApiUsingPermissions) // 获取jssdk调用权限
		wechatRouter.GET("userInfo", wechatApi.GetSnsapiUserInfo)     // 获取(公众号)微信用户信息
	}
}
