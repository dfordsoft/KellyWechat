package models

import (
	"github.com/astaxie/beego"
)

func FacialMask(req *Request, resp *Response) error {
	shopId, _ := beego.AppConfig.Int("mm_shop_id")
	return itemList(shopId, req, resp)
}
