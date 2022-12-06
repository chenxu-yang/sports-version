package service

import (
	"wxcloudrun-golang/app/auth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func WeChatLogin(c *gin.Context) {
	code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := auth.WXLogin(code)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	// 保存登录态
	session := sessions.Default(c)
	session.Set("openid", wxLoginResp.OpenId)
	session.Set("sessionKey", wxLoginResp.SessionKey)
	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
	mySession := auth.GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
	c.String(200, mySession)
}
