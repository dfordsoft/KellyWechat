package models

var (
	AboutContent = `您好，感谢关注衣丽社区微信公众号i-yiili，希望我们能为您的生活提供一点帮助。
            我们会通过微信公众号向您不定期地推送一些涉及美容美发化妆、穿着搭配、运动健康、饮食营养、情绪修养等跟女性生活息息相关的文章。
            同时借助微信公众号强大的可定制性，我们也提供了一系列丰富的功能，并且将不断增加新的实用的功能。
            请输入help查看使用帮助。`
	HelpContent = `
            输入“微店”或“wd”，显示衣丽社区关联的微店地址。
            输入“女装”或“nz”，显示衣丽社区主持出售的女装商品信息。
            输入“面膜”或“mm”，显示衣丽社区主持出售的面膜商品信息。
            输入商品id（如：310148677），显示衣丽社区主持出售的对应商品信息。
            输入任意关键字，显示包含该关键字的衣丽社区主持出售的商品信息。
            其他更多功能，敬请期待。`
)

func About(req *Request, resp *Response) error {
	var a WXMPItem
	resp.MsgType = News
	resp.ArticleCount = 1
	a.Description = AboutContent
	a.Title = `关于衣丽社区微信公众号`
	a.PicUrl = `https://dn-kelly.qbox.me/upload/img/mc5y9.full.jpg`
	a.Url = "https://yii.li/post/19"
	resp.Articles = append(resp.Articles, &a)
	resp.FuncFlag = 1

	return nil
}

func Help(req *Request, resp *Response) error {
	var a WXMPItem
	resp.MsgType = News
	resp.ArticleCount = 1
	a.Description = HelpContent
	a.Title = `衣丽微信公众号使用帮助`
	a.PicUrl = "https://dn-kelly.qbox.me/upload/img/zb5y9.full.jpg"
	a.Url = `https://yii.li/post/20`
	resp.Articles = append(resp.Articles, &a)
	resp.FuncFlag = 1

	return nil
}
