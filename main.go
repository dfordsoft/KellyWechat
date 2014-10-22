package main

import (
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/controllers"
)

func main() {
	wxmp := new(controllers.WXMPController)
	beego.Router("/", wxmp, "get:Get;post:Post")
	beego.Router("/wd/add", wxmp, "post:Post")
	beego.Router("/article/add", wxmp, "post:Post")
	beego.HttpPort = 8091
	beego.Run()
}
