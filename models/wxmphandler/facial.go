package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models/wd"
)

func FacialMask(req *Request, resp *Response) error {
	shopId, _ := beego.AppConfig.Int("mm_shop_id")
	qs := models.Items()
	items := make([]models.WDItem, 10)
	n, err := qs.Limit(10).Filter("shop_id", shopId).All(&items)
	if err != nil || n == 0 {
		resp.Content = `没开店哦:(`
		return nil
	}

	resp.MsgType = News
	resp.ArticleCount = int(n)
	a := make([]WXMPItem, n)
	for i, item := range items {
		if n > 6 && i >= 5 {
			resp.ArticleCount = 5
			break
		}
		a[i].Description = ``
		a[i].Title = item.Name
		a[i].PicUrl = item.Logo
		a[i].Url = fmt.Sprintf(`http://wd.koudai.com/i/%d`, item.Uuid)
		resp.Articles = append(resp.Articles, &a[i])
	}

	if n > 6 {
		wdShop := &models.WDShop{}
		wdShop.Id = shopId
		if wdShop.Get("id") == nil {
			resp.ArticleCount++
			shopItem := &WXMPItem{}
			shopItem.Description = wdShop.Note
			shopItem.Title = `宝贝数量较多，请进入微店查看更多 - ` + wdShop.Name
			shopItem.PicUrl = wdShop.Logo
			shopItem.Url = fmt.Sprintf(`http://wd.koudai.com/s/%d`, wdShop.Uuid)
			resp.Articles = append(resp.Articles, shopItem)
		}
	}

	resp.FuncFlag = 1
	return nil
}
