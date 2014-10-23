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

func (s *WDShop) Insert() error {
	if _, err := orm.NewOrm().Insert(s); err != nil {
		return err
	}
	return nil
}

func (s *WDShop) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

func (s *WDShop) Get(fields ...string) error {
	if err := orm.NewOrm().Read(s, fields...); err != nil {
		return err
	}
	return nil
}

func Shops() orm.QuerySeter {
	o := orm.NewOrm()

	qs := o.QueryTable("w_d_shop").OrderBy("-Id")
	return qs
}
