package routers

import (
	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/aria2-http-api/controllers"
)

func init() {
	api := beego.NewNamespace("/aria2",
		beego.NSNamespace("/v1",
			beego.NSRouter("/addUrl", &controllers.Aria2Controller{}, "post:AddUrl"),
			beego.NSRouter("/tellStatus", &controllers.Aria2Controller{}, "get:TellStatus"),
			beego.NSRouter("/tellActive", &controllers.Aria2Controller{}, "get:TellActive"),
			beego.NSRouter("/remove", &controllers.Aria2Controller{}, "delete:Remove"),
			beego.NSRouter("/pause", &controllers.Aria2Controller{}, "put:Pause"),
			beego.NSRouter("/unpause", &controllers.Aria2Controller{}, "put:UnPause"),
			beego.NSRouter("/pauseAll", &controllers.Aria2Controller{}, "put:PauseAll"),
		),
	)
	beego.AddNamespace(api)
}
