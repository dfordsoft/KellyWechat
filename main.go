package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/missdeer/KellyWechat/controllers"
	"github.com/missdeer/KellyWechat/models/wd"
)

func main() {

	driverName := beego.AppConfig.String("driver_name")
	dataSource := beego.AppConfig.String("data_source")
	maxIdle, _ := beego.AppConfig.Int("max_idle_conn")
	maxOpen, _ := beego.AppConfig.Int("max_open_conn")

	models.Init()

	// set default database
	err := orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxOpen)
	if err != nil {
		beego.Error(err)
	}
	orm.RunCommand()

	err = orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Error(err)
	}

	wxmp := new(controllers.WXMPController)
	beego.Router("/", wxmp, "get:Get;post:Post")
	wxmp.GetAccessToken()
	wxmp.SetupMenu()
	go wxmp.UpdateAccessToken()

	wd := new(controllers.WDController)
	beego.Router("/wd/add/:id", wd, "post:SubmitWD")

	beego.Router("/article/add", wxmp, "post:Post")

	beego.HttpPort, _ = beego.AppConfig.Int("http_port")
	beego.Run()
}
