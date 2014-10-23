package models

import (
	"fmt"
	"github.com/missdeer/KellyWechat/models/wd"
)

func WeiDian(req *Request, resp *Response) error {
	qs := models.Shops()
	shops := make([]models.WDShop, 5)
	n, err := qs.Limit(5).All(&shops)
	if err != nil || n == 0 {
		resp.Content = `没开店哦:(`
		return nil
	}

	resp.MsgType = News
	resp.ArticleCount = int(n)
	a := make([]WXMPItem, n)
	for i, shop := range shops {
		a[i].Description = ``
		a[i].Title = shop.Name + " - " + shop.Note
		a[i].PicUrl = shop.Logo
		a[i].Url = fmt.Sprintf(`http://wd.koudai.com/s/%d`, shop.Uuid)
		resp.Articles = append(resp.Articles, &a[i])
	}
	resp.FuncFlag = 1
	return nil
}
