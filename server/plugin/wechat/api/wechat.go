package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/common"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/wechat/pkg/tools"
	"github.com/gin-gonic/gin"
	"time"
)

type WechatApi struct{}

// GetJsApiUsingPermissions 获取用于调用微信JS接口的使用权限
// @Tags WeChat
// @Summary 微信JS接口的使用权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param url query string true "url"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wechat/jsapi [get]
func (cuApi *WechatApi) GetJsApiUsingPermissions(c *gin.Context) {
	var url string
	url = c.Query("url")
	if url == "" {
		response.FailWithMessage("empty url", c)
		return
	}

	noncestr := tools.NewRandGenerator(tools.WithLength(12)).Generate()
	jsapi_ticket := common.GetWechatPublicJsApiTicket()
	data := tools.GetJsApiUsingPermissions{
		Noncestr:    noncestr,
		JsapiTicket: jsapi_ticket,
		Timestamp:   time.Now().Unix(),
		Url:         url,
	}

	sha1 := data.Sha1()
	appId := global.GlobalConfig.Wechat.PubWxConfig.AppID // 公众号appId

	result := tools.GetJsApiUsingPermissionsResp{
		GetJsApiUsingPermissions: data,
		AppId:                    appId,
		Signature:                sha1,
	}

	response.OkWithData(result, c)
}

// GetSnsapiUserInfo 获取(公众号)微信用户信息
// @Tags WeChat
// @Summary 获取(公众号)微信用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param code query string true "用户同意授权，获取code"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wechat/userInfo [get]
func (cuApi *WechatApi) GetSnsapiUserInfo(c *gin.Context) {
	// 获取微信配置
	wechatConfig := global.GlobalConfig.Wechat

	// 获取code
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("empty code", c)
		return
	}

	// 获取web access token
	result, err := wechatConfig.PubWxConfig.GetWebAccessToken(code)
	if err != nil {
		response.FailWithMessage("get access token failed", c)
		return
	}
	if result.ErrCode != 0 {
		response.FailWithMessage(result.ErrMsg, c)
		return
	}

	// 获取用户信息
	userInfo, err := tools.GetSnsapiUserInfo(result.AccessToken, result.OpenId)
	if err != nil {
		response.FailWithMessage("get user info failed", c)
		return
	}

	response.OkWithData(userInfo, c)
}
