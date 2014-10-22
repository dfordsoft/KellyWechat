package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	TOKEN       = "yiiliwechattoken"
	Text        = "text"
	Location    = "location"
	Image       = "image"
	Link        = "link"
	Event       = "event"
	Music       = "music"
	News        = "news"
	HelpContent = `您好，感谢关注衣丽社区微信公众号i-yiili，希望我们能为您的生活提供一点帮助。
            我们会通过微信公众号向您不定期地推送一些涉及美容美发化妆、穿着搭配、运动健康、饮食营养、情绪修养等跟女性生活息息相关的文章。
            同时借助微信公众号强大的可定制性，我们也提供了一系列丰富的功能，并且将不断增加新的实用的功能。
            输入“文章+日期”或“wz+日期”如（文章20141022 或 wz20141022），显示该天衣丽社区微信公众号发布的文章链接。
            输入“微店”或“wd”，显示衣丽社区关联的微店地址。
            输入“女装”或“nz”，显示衣丽社区主持出售的女装商品信息。
            输入“面膜”或“mm”，显示衣丽社区主持出售的面膜商品信息。
            输入商品id（如：310148677），显示衣丽社区主持出售的对应商品信息。
            输入任意关键字，显示包含该关键字的衣丽社区主持出售的商品信息。
            其他更多功能，敬请期待。`
)

type msgBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

type Request struct {
	XMLName                xml.Name `xml:"xml"`
	msgBase                         // base struct
	Location_X, Location_Y float32
	Scale                  int
	Label                  string
	PicUrl                 string
	MsgId                  int
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	msgBase
	ArticleCount int     `xml:",omitempty"`
	Articles     []*item `xml:"Articles>item,omitempty"`
	FuncFlag     int
}

type item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	signature := this.Input().Get("signature")
	beego.Info("signature:" + signature)
	timestamp := this.Input().Get("timestamp")
	beego.Info("timestamp:" + timestamp)
	nonce := this.Input().Get("nonce")
	beego.Info("nonce:" + nonce)
	echostr := this.Input().Get("echostr")
	beego.Info("echostr:" + echostr)
	beego.Info(Signature(timestamp, nonce))
	if Signature(timestamp, nonce) == signature {
		beego.Info("signature matched")
		this.Ctx.WriteString(echostr)
	} else {
		beego.Info("signature not matched")
		this.Ctx.WriteString("")
	}
}

func (this *MainController) Post() {
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	beego.Info(string(body))
	var wreq *Request
	if wreq, err = DecodeRequest(body); err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	beego.Info(wreq.Content)
	wresp, err := dealwith(wreq)
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	data, err := wresp.Encode()
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	this.Ctx.WriteString(string(data))
	return
}

func dealwith(req *Request) (resp *Response, err error) {
	resp = NewResponse()
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Text
	beego.Info(req.MsgType)
	beego.Info(req.Content)
	if req.MsgType == Text {
		userInputText := strings.Trim(strings.ToLower(req.Content), " ")
		if userInputText == "help" || userInputText == `帮助` {
			resp.Content = HelpContent
			return resp, nil
		}
		if userInputText == "wd" || userInputText == `微店` {
			resp.MsgType = News
			resp.ArticleCount = 2
			var a item
			a.Description = `女装专卖，经典时尚大方，赶快来看看吧:)`
			a.Title = `凯莉小姐的梦想女装店，经典时尚大方，适合都市年轻女性，赶快来吧:)`
			a.PicUrl = "http://wd.geilicdn.com/vshop215091300-1413902752.jpg"
			a.Url = "http://shopwd.yii.li"
			resp.Articles = append(resp.Articles, &a)
			var a2 item
			a2.Description = `多款精品面膜，效果超好，输入mm看看吧:)`
			a2.Title = `清韵诗面膜，多款精品面膜，效果超好，输入mm看看吧:)`
			a2.PicUrl = "http://wd.geilicdn.com/vshop1015143-1413815285.jpg"
			a2.Url = "http://wd.koudai.com/s/1015143"
			resp.Articles = append(resp.Articles, &a2)

			resp.FuncFlag = 1
			return resp, nil
		}
		if userInputText == "mm" || userInputText == `面膜` {
			resp.MsgType = News
			resp.ArticleCount = 6
			var a [6]item
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
			return resp, nil
		}

		strs := strings.Split(req.Content, ".")
		var resurl string
		var a item
		if len(strs) == 1 {
			resurl = "https://raw.github.com/astaxie/gopkg/master/" + strings.Trim(strings.ToLower(strs[0]), " ") + "/README.md"
			a.Url = "https://github.com/astaxie/gopkg/tree/master/" + strings.Trim(strings.ToLower(strs[0]), " ") + "/README.md"
		} else {
			var other []string
			for k, v := range strs {
				if k < (len(strs) - 1) {
					other = append(other, strings.Trim(strings.ToLower(v), " "))
				} else {
					other = append(other, strings.Trim(strings.Title(v), " "))
				}
			}
			resurl = "https://raw.github.com/astaxie/gopkg/master/" + strings.Join(other, "/") + ".md"
			a.Url = "https://github.com/astaxie/gopkg/tree/master/" + strings.Join(other, "/") + ".md"
		}
		beego.Info(resurl)
		rsp, err := http.Get(resurl)
		if err != nil {
			resp.Content = "不存在该包内容"
			return nil, err
		}
		defer rsp.Body.Close()
		if rsp.StatusCode == 404 {
			resp.Content = "找不到你要查询的包:" + req.Content
			return resp, nil
		}
		resp.MsgType = News
		resp.ArticleCount = 1
		body, err := ioutil.ReadAll(rsp.Body)
		beego.Info(string(body))
		a.Description = string(body)
		a.Title = req.Content
		a.PicUrl = "http://bbs.gocn.im/static/image/common/logo.png"
		resp.Articles = append(resp.Articles, &a)
		resp.FuncFlag = 1
	} else {
		resp.Content = "暂时还不支持其他的类型"
	}
	return resp, nil
}

func Signature(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func DecodeRequest(data []byte) (req *Request, err error) {
	req = &Request{}
	if err = xml.Unmarshal(data, req); err != nil {
		return
	}
	req.CreateTime *= time.Second
	return
}

func NewResponse() (resp *Response) {
	resp = &Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	return
}

func (resp Response) Encode() (data []byte, err error) {
	resp.CreateTime = time.Second
	data, err = xml.Marshal(resp)
	return
}
