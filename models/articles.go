package models

func Articles(req *Request, resp *Response) error {
	resp.MsgType = News
	resp.ArticleCount = 6
	var a [6]Item
	a[0].Description = `售价仅88元一盒，赶快来看看吧:)`
	a[0].Title = `清韵诗完美净肤面膜，售价仅88元一盒，赶快来看看吧:)`
	a[0].PicUrl = "http://wd.geilicdn.com/vshop1015143-1413814758-1.jpg"
	a[0].Url = "http://wd.koudai.com/item.html?itemID=307973157"
	resp.Articles = append(resp.Articles, &a[0])
	a[1].Description = `售价仅88元一盒，赶快来看看吧:)`
	a[1].Title = `清韵诗焕颜靓白面膜，售价仅88元一盒，赶快来看看吧:)`
	a[1].PicUrl = "http://wd.geilicdn.com/vshop1015143-1413814687-1.jpg"
	a[1].Url = "http://wd.koudai.com/item.html?itemID=307970560"
	resp.Articles = append(resp.Articles, &a[1])
	a[2].Description = `售价仅88元一盒，赶快来看看吧:)`
	a[2].Title = `清韵诗弹力胶原面膜，售价仅88元一盒，赶快来看看吧:)`
	a[2].PicUrl = "http://wd.geilicdn.com/vshop1015143-1413814573-1.jpg"
	a[2].Url = "http://wd.koudai.com/item.html?itemID=307966427"
	resp.Articles = append(resp.Articles, &a[2])
	a[3].Description = `售价仅88元一盒，赶快来看看吧:)`
	a[3].Title = `清韵诗植物修复面膜，售价仅88元一盒，赶快来看看吧:)`
	a[3].PicUrl = "http://wd.geilicdn.com/vshop1015143-1413814512-1.jpg"
	a[3].Url = "http://wd.koudai.com/item.html?itemID=307964195"
	resp.Articles = append(resp.Articles, &a[3])
	a[4].Description = `售价仅88元一盒，赶快来看看吧:)`
	a[4].Title = `清韵诗魔法冰瓷面膜，售价仅88元一盒，赶快来看看吧:)`
	a[4].PicUrl = "http://wd.geilicdn.com/vshop1015143-1413814406-1.jpg"
	a[4].Url = "http://wd.koudai.com/item.html?itemID=307960310"
	resp.Articles = append(resp.Articles, &a[4])
	a[5].Description = `售价仅88元一盒，赶快来看看吧:)`
	a[5].Title = `清韵诗水嫩宝贝面膜，售价仅88元一盒，赶快来看看吧:)`
	a[5].PicUrl = "http://wd.geilicdn.com/vshop1015143-1413814179-1.jpg"
	a[5].Url = "http://wd.koudai.com/item.html?itemID=307953149"
	resp.Articles = append(resp.Articles, &a[5])

	resp.FuncFlag = 1
	return nil
}
