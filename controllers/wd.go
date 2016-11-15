package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/missdeer/KellyWechat/models/wd"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

type WDResponseItemInfo struct {
	ItemId   uint64
	ItemName string
	Img      string
	Price    string
}

type WDResponseItemList struct {
	Status WDResponseStatus
	Result []WDResponseItemInfo
}

type WDController struct {
	beego.Controller
}

func (this *WDController) getItemList(wdShop *models.WDShop, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		beego.Error("read response error: ", err)
		this.Data["json"] = map[string]string{"error": "reading item list response error"}
		this.ServeJSON()
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var itemList WDResponseItemList
	if err = json.Unmarshal(body, &itemList); err != nil {
		beego.Error("unmarshalling item list error: ", err)
		this.Data["json"] = map[string]string{"error": "unmarshalling item list response error"}
		this.ServeJSON()
		return err
	}
	fmt.Println("item list: ", itemList)

	if len(itemList.Result) == 0 {
		beego.Error("empty item list")
		return errors.New("empty item list")
	}

	for _, item := range itemList.Result {
		wdItem := &models.WDItem{}
		wdItem.Uuid = item.ItemId

		if wdItem.Get("uuid") != nil {
			endPos := strings.Index(item.Img, "?")
			wdItem.Logo = item.Img[:endPos]
			wdItem.Name = item.ItemName
			wdItem.Shop = wdShop

			beego.Info("do insert item record")
			wdItem.Insert()
		} else {
			endPos := strings.Index(item.Img, "?")
			wdItem.Logo = item.Img[:endPos]
			wdItem.Name = item.ItemName
			wdItem.Shop = wdShop

			beego.Info("do update item record")
			wdItem.Update("id")
		}
	}

	return nil
}

func (this *WDController) SubmitWD() {
	id := this.GetString(":id")
	// get shop info
	// http://wd.koudai.com/wd/shop/getPubInfo?param={"userID":215091300,"f_seller_id":""}
	shopInfoUrl := fmt.Sprintf(`http://wd.koudai.com/wd/shop/getPubInfo?param={"userID":%s,"f_seller_id":""}`, id)

	resp, err := http.Get(shopInfoUrl)
	if err != nil {
		beego.Error("read response error: ", err)
		this.Data["json"] = map[string]string{"error": "reading shop info response error"}
		this.ServeJSON()
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
		this.ServeJSON()
		return
	}
	wdShop := &models.WDShop{}
	wdShop.Uuid, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		beego.Error("read response error: ", err)
		this.Data["json"] = map[string]string{"error": "reading response error"}
		this.ServeJSON()
		return
	}

	if wdShop.Get("uuid") != nil {
		endPos := strings.Index(shopInfo.Result.Logo, "?")
		wdShop.Logo = shopInfo.Result.Logo[:endPos]
		wdShop.Name = shopInfo.Result.ShopName
		wdShop.Note = shopInfo.Result.Note
		beego.Info("do insert shop record")
		wdShop.Insert()
	} else {
		endPos := strings.Index(shopInfo.Result.Logo, "?")
		wdShop.Logo = shopInfo.Result.Logo[:endPos]
		wdShop.Name = shopInfo.Result.ShopName
		wdShop.Note = shopInfo.Result.Note
		beego.Info("do update shop record")
		wdShop.Update("id")
	}

	// get item list
	// http://wd.koudai.com/wd/item/getIsTopList?param={"userid":215091300,"pageNum":0,"pageSize":49,"isTop":0,"f_seller_id":""}
	for i := 0; ; i++ {
		itemListUrl := fmt.Sprintf(`http://wd.koudai.com/wd/item/getIsTopList?param={"userid":%s,"pageNum":%d,"pageSize":49,"isTop":0,"f_seller_id":""}`, id, i)
		if err = this.getItemList(wdShop, itemListUrl); err != nil {
			break
		}
	}

	// get item detail
	// http://wd.koudai.com/wd/item/getPubInfo?param={"itemID":310148677,"page":1}

	this.Data["json"] = map[string]string{"ok": "200"}
	this.ServeJSON()
}
