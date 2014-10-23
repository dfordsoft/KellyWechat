package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

type WDResponseStatus struct {
	StatusCode   int
	StatusReason string
}

type WDResponseShopResult struct {
	ShopName string
	Logo     string
	Note     string
}

type WDResponseShopInfo struct {
	Status WDResponseStatus
	Result WDResponseShopResult
}

type WDController struct {
	beego.Controller
}

func (this *WDController) SubmitWD() {
	id := this.GetString(":id")
	// get shop info
	// http://wd.koudai.com/wd/shop/getPubInfo?param={"userID":215091300,"f_seller_id":""}
	shopInfoUrl := fmt.Sprintf(`http://wd.koudai.com/wd/shop/getPubInfo?param={"userID":%s,"f_seller_id":""}`, id)

	resp, err := http.Get(shopInfoUrl)
	if err != nil {
		beego.Error("read response error: ", err)
		this.Data["json"] = map[string]string{"error": "reading response error"}
		this.ServeJson()
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var shopInfo WDResponseShopInfo
	json.Unmarshal(body, &shopInfo)
	fmt.Println("shop info: ", shopInfo)
	if shopInfo.Status.StatusCode != 0 ||
		len(shopInfo.Result.ShopName) == 0 ||
		len(shopInfo.Result.Logo) == 0 {
		this.Data["json"] = map[string]string{"error": "seemly an invalid shop"}
		this.ServeJson()
		return
	}
	// get item list
	// http://wd.koudai.com/wd/item/getIsTopList?param={"userid":215091300,"pageNum":0,"pageSize":49,"isTop":0,"f_seller_id":""}

	// get item detail
	// http://wd.koudai.com/wd/item/getPubInfo?param={"itemID":310148677,"page":1}

	this.Data["json"] = map[string]string{"ok": "200"}
	this.ServeJson()
}
