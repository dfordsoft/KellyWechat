package models

import (
	"strings"
)

func ItemId(req *Request, resp *Response) error {
	resp.MsgType = News
	resp.ArticleCount = 1
	var a Item
	a.Description = `女装专卖，经典时尚大方，赶快来看看吧:)`
	a.Title = `凯莉小姐的梦想女装店，经典时尚大方，适合都市年轻女性，赶快来吧:)`
	a.PicUrl = "http://wd.geilicdn.com/vshop215091300-1413902752.jpg"
	userInputText := strings.Trim(strings.ToLower(req.Content), " ")
	a.Url = "http://wd.koudai.com/item.html?itemID=" + userInputText
	resp.Articles = append(resp.Articles, &a)
	resp.FuncFlag = 1
	return nil
}
