package controllers

import (
	"github.com/astaxie/beego"
)

const (
	aria2HttpTokenName = "ARIA2_HTTP_API_TOKEN"
)

var (
	aria2HTTPAPIToken = beego.AppConfig.String("security::aria2HTTPAPIToken")
)

type Data map[string]interface{}

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

	// 获取 头部信息
	tokenHeader := b.Ctx.Input.Header(aria2HttpTokenName)
	if tokenHeader != aria2HTTPAPIToken {
		b.JsonError("token auth", "token auth error", "")
		b.StopRun()
	}

}
