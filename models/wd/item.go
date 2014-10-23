package models

type WDItem struct {
	Id   int
	Shop *WDShop `orm:"rel(fk)"`
	Uuid uint64  `orm:"index"`
	Logo string  `orm:"type(text)"`
	Name string  `orm:"size(64);unique"`
	Note string  `orm:"type(text)"`
}
