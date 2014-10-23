package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models/wd"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ItemId(req *Request, resp *Response) error {
	userInputText := strings.Trim(strings.ToLower(req.Content), " ")
	uuid, err := strconv.ParseUint(userInputText, 10, 64)
	if err != nil {
		beego.Error("incorrect input ", userInputText)
		return nil
	}
	item := &models.WDItem{}
	item.Uuid = uuid
	if item.Get("uuid") != nil {
		beego.Error("not found ", userInputText)
		return nil
	}

	var a WXMPItem
	resp.MsgType = News
	resp.ArticleCount = 1
	a.Description = `点击查看详细信息哦:)`
	a.Title = item.Name
	a.PicUrl = item.Logo
	a.Url = fmt.Sprintf(`http://wd.koudai.com/i/%d`, item.Uuid)
	resp.Articles = append(resp.Articles, &a)
	resp.FuncFlag = 1
	return nil
}

func composeItemListReponse(items []models.WDItem, shopId int, req *Request, resp *Response) error {
	n := len(items)
	resp.MsgType = News
	resp.ArticleCount = n
	arrayLength := n
	if arrayLength < 5 {
		arrayLength++
	} else {
		arrayLength = 6
	}
	a := make([]WXMPItem, arrayLength)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 5; i++ {
		if n > 6 && i >= 5 {
			resp.ArticleCount = 5
			break
		}
		item := items[rand.Intn(n)]
		a[i].Description = ``
		a[i].Title = item.Name
		a[i].PicUrl = item.Logo
		a[i].Url = fmt.Sprintf(`http://wd.koudai.com/i/%d`, item.Uuid)
		resp.Articles = append(resp.Articles, &a[i])
	}

	if n > 6 && shopId >= 0 {
		wdShop := &models.WDShop{}
		wdShop.Id = shopId
		shopItem := &WXMPItem{}
		if wdShop.Get("id") == nil {
			resp.ArticleCount++
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

func itemList(shopId int, req *Request, resp *Response) error {
	qs := models.Items()
	var items []models.WDItem
	n, err := qs.Limit(500).Filter("shop_id", shopId).All(&items)
	if err != nil || n == 0 {
		resp.Content = `没开这样的店哦:(`
		return nil
	}

	return composeItemListReponse(items, shopId, req, resp)
}

func SearchItems(req *Request, resp *Response) error {
	userInputText := strings.Trim(strings.ToLower(req.Content), " ")
	qs := models.Items()
	var items []models.WDItem
	n, err := qs.Limit(500).Filter("name__icontains", userInputText).All(&items)
	if err != nil || n == 0 {
		resp.Content = `查不到包含关键字“` + userInputText + `”的宝贝哦:(`
		return nil
	}

	return composeItemListReponse(items, -1, req, resp)
}
