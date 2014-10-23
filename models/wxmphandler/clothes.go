package models

import (
	"github.com/astaxie/beego"
)

func Clothes(req *Request, resp *Response) error {
	shopId, _ := beego.AppConfig.Int("yf_shop_id")
	return itemList(shopId, req, resp)
}
