package models

func WeiDian(req *Request, resp *Response) error {
	resp.MsgType = News
	resp.ArticleCount = 2
	var a Item
	a.Description = `女装专卖，经典时尚大方，赶快来看看吧:)`
	a.Title = `凯莉小姐的梦想女装店，经典时尚大方，适合都市年轻女性，赶快来吧:)`
	a.PicUrl = "http://wd.geilicdn.com/vshop215091300-1413902752.jpg"
	a.Url = "http://shopwd.yii.li"
	resp.Articles = append(resp.Articles, &a)
	var a2 Item
	a2.Description = `多款精品面膜，效果超好，输入mm看看吧:)`
	a2.Title = `清韵诗面膜，多款精品面膜，效果超好，输入mm看看吧:)`
	a2.PicUrl = "http://wd.geilicdn.com/vshop1015143-1413815285.jpg"
	a2.Url = "http://wd.koudai.com/s/1015143"
	resp.Articles = append(resp.Articles, &a2)

	resp.FuncFlag = 1
	return nil
}
