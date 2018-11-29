package main

import (
	"github.com/astaxie/beego"
	_ "github.com/chanyipiaomiao/aria2-http-api/routers"
)

func main() {
	beego.Run()
}
