package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models/wd"
)

func Clothes(req *Request, resp *Response) error {
	shopId, _ := beego.AppConfig.Int("yf_shop_id")
	qs := models.Items()
	items := make([]models.WDItem, 10)
	n, err := qs.Limit(10).Filter("ShopId", shopId).All(&items)
	if err != nil || n == 0 {
		resp.Content = `没开店哦:(`
		return nil
	}

	resp.MsgType = News
	resp.ArticleCount = int(n)
	a := make([]WXMPItem, n)
	for i, item := range items {
		a[i].Description = ``
		a[i].Title = item.Name
		a[i].PicUrl = item.Logo
		a[i].Url = fmt.Sprintf(`http://wd.koudai.com/i/%d`, item.Uuid)
		resp.Articles = append(resp.Articles, &a[i])
	}
	resp.FuncFlag = 1
	return nil
}
