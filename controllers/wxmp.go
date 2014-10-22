package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	TOKEN = "yiiliwechattoken"
)

type WXMPController struct {
	beego.Controller
}

func (this *WXMPController) Get() {
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

func (this *WXMPController) Post() {
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	beego.Info(string(body))
	var wreq *models.Request
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

func dealwith(req *models.Request) (resp *models.Response, err error) {
	resp = NewResponse()
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = models.Text
	beego.Info(req.MsgType)
	beego.Info(req.Content)
	if req.MsgType == models.Text {
		userInputText := strings.Trim(strings.ToLower(req.Content), " ")
		switch userInputText {
		case "help", `帮助`:
			models.Help(req, resp)
		case "wd", `微店`:
			models.WeiDian(req, resp)
		case "mm", `面膜`:
			models.FacialMask(req, resp)
		case "nv", "yf", "yifu", `女装`, `衣服`:
			models.Clothes(req, resp)
		default:
			matched, err := regexp.MatchString("[0-9]+", userInputText)
			if err == nil && matched == true {
				models.ItemId(req, resp)
				break
			}
			matched, err = regexp.MatchString("(wz|article|文章)[0-9]{8,8}", userInputText)
			if err == nil && matched == true {
				models.Articles(req, resp)
				break
			}
			resp.Content = "衣丽已经很努力地在学习了，但仍然不能理解您的需求，请您输入help查看衣丽能懂的一些命令吧:("
		}
	} else {
		resp.Content = "暂时还不支持其他的命令类型，请输入help查看说明。"
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

func DecodeRequest(data []byte) (req *models.Request, err error) {
	req = &models.Request{}
	if err = xml.Unmarshal(data, req); err != nil {
		return
	}
	req.CreateTime *= time.Second
	return
}

func NewResponse() (resp *models.Response) {
	resp = &models.Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	return
}
