package controllers

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models/wxmphandler"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	access_token string
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint   `json:"expires_in"`
}

type WXMPController struct {
	beego.Controller
}

func (this *WXMPController) UpdateAccessToken() {
	timer := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-timer.C:
			if this.GetAccessToken() != nil {
				beego.Error("updating access token failed, try again")
				this.GetAccessToken()
			}
		}
	}
	timer.Stop()
}

func (this *WXMPController) GetAccessToken() error {
	appId := beego.AppConfig.String("appid")
	appSecret := beego.AppConfig.String("appsecret")
	url := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, appId, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		beego.Error("read response error: ", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var atr AccessTokenResponse
	if err = json.Unmarshal(body, &atr); err != nil {
		beego.Error("unmarshalling get access token response error: ", err)
		return err
	}
	access_token = atr.AccessToken
	beego.Info("get access_token: ", access_token)
	return nil
}

type SetupMenuResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (this *WXMPController) SetupMenu() error {
	url := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s`, access_token)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(models.MenuDefine))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		beego.Error("post setup menu command failed: ", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("reading setup menu response failed: ", err)
		return err
	}

	var r SetupMenuResponse
	if err = json.Unmarshal(body, &r); err != nil {
		beego.Error("unmarshalling setup menu response error: ", err)
		return err
	}

	if r.Errcode != 0 {
		beego.Error("setup menu failed: ",r.Errcode, r.Errmsg)
		return errors.New(r.Errmsg)
	}

	beego.Info("setup menu successfully")
	return nil
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
		case "nz", "yf", "yifu", `女装`, `衣服`:
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
			models.SearchItems(req, resp)
		}
	} else {
		models.Help(req, resp)
	}
	return resp, nil
}

func Signature(timestamp, nonce string) string {
	wxmpToken := beego.AppConfig.String("wxmp_token")
	strs := sort.StringSlice{wxmpToken, timestamp, nonce}
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
