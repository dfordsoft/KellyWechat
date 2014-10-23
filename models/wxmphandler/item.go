package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models/wd"
	"strconv"
	"strings"
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
