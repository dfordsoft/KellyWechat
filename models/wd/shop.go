package models

import (
	"github.com/astaxie/beego/orm"
)

type WDShop struct {
	Id   int
	Uuid uint64 `orm:"index"`
	Logo string `orm:"type(text)"`
	Name string `orm:"size(64);unique"`
	Note string `orm:"type(text)"`
}

func Init() {
	orm.RegisterModel(new(WDShop), new(WDItem))
}
