package main

import (
    "github.com/astaxie/beego"
    "github.com/missdeer/KellyWechat/controllers"
)

func main() {

    r := new(controllers.MainController)
    beego.Router("/", r, "get:Get;post:Post")
    beego.HttpPort = 8091
    beego.Run()
}
