package controllers

import "github.com/astaxie/beego"

const (
	aria2HttpTokenName = "ARIA2_HTTP_API_TOKEN"
)

var (
	aria2Token = beego.AppConfig.String("security::aria2Token")
)

type BaseController struct {
	beego.Controller
}

func (b *BaseController) json(entryType, errmsg string, statuscode int, data interface{}) {
	msg := map[string]interface{}{
		"entryType":  entryType,
		"error":      errmsg,
		"statusCode": statuscode,
		"data":       data,
	}
	b.Data["json"] = msg
	b.ServeJSON()
}

func (b *BaseController) JsonError(entryType, errmsg string, data interface{}) {
	b.json(entryType, errmsg, 1, data)
}

func (b *BaseController) JsonOK(entryType string, data interface{}) {
	b.json(entryType, "", 0, data)
}

func (b *BaseController) Prepare() {

	/// 获取客户端IP
	b.Data["RemoteIP"] = b.Ctx.Input.IP()

	// 获取 头部信息
	token := b.Ctx.Input.Header(aria2HttpTokenName)
	if token != aria2Token {
		b.JsonError("token auth", "need token header", "")
		b.StopRun()
	}
}
