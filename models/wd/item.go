package models

import (
	"github.com/astaxie/beego/orm"
)

type WDItem struct {
	Id   int
	Shop *WDShop `orm:"rel(fk)"`
	Uuid uint64  `orm:"index"`
	Logo string  `orm:"type(text)"`
	Name string  `orm:"size(64);unique"`
}

func (i *WDItem) Insert() error {
	if _, err := orm.NewOrm().Insert(i); err != nil {
		return err
	}
	return nil
}

func (i *WDItem) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(i, fields...); err != nil {
		return err
	}
	return nil
}

func (i *WDItem) Get(fields ...string) error {
	if err := orm.NewOrm().Read(i, fields...); err != nil {
		return err
	}
	return nil
}

func Items() orm.QuerySeter {
	o := orm.NewOrm()

	qs := o.QueryTable("w_d_item").OrderBy("-Id")
	return qs
}
