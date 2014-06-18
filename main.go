package main

import (
    "log"
    "net/http"
    "github.com/sidbusy/weixinmp"
)

func main() {
    // 注册处理函数
    http.HandleFunc("/receiver", receiver)
    log.Fatal(http.ListenAndServe(":8091", nil))
}

func receiver(w http.ResponseWriter, r *http.Request) {
    token := "" // 微信公众平台的Token
    appid := "" // 微信公众平台的AppID
    secret := "" // 微信公众平台的AppSecret
    // 仅被动响应消息时可不填写appid、secret
    // 仅主动发送消息时可不填写token
    mp := weixinmp.New(token, appid, secret)
    // 检查请求是否有效
    // 仅主动发送消息时不用检查
    if !mp.Request.IsValid(w, r) {
        log.Println("invalid request")
        return
    }
    // 判断消息类型
    if mp.Request.MsgType == weixinmp.MsgTypeText {
        // 回复消息
        mp.ReplyTextMsg(w, "Hello, 世界")
    }
}
