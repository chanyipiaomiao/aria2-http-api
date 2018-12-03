package routers

import (
	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/aria2-http-api/controllers"
)

func init() {
	api := beego.NewNamespace("/aria2",
		beego.NSNamespace("/v1",
			beego.NSRouter("/addUrl", &controllers.Aria2Controller{}),
		),
	)
	beego.AddNamespace(api)
}
