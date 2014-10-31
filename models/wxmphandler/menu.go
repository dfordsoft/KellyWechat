package models

type ClickItem struct {
	Type string
	Name string
	Key  string
}

type ViewItem struct {
	Type string
	Name string
	Url  string
}

var (
	MenuDefine = []byte(`
    {
        "button":[
        {
            "name":"微店",
            "sub_button":[
            {  
                "type":"click",
                "name":"微店列表",
                "key":"V1001_TODAY_LISTWD"
            },
            {  
                "type":"click",
                "name":"面膜列表",
                "key":"V1001_TODAY_LISTMM"
            },
            {
                "type":"view",
                "name":"面膜微店",
                "url":"http://v.qq.com/"
            },
            {  
                "type":"click",
                "name":"服装列表",
                "key":"V1001_TODAY_LISTYF"
            },
            {
                "type":"view",
                "name":"服装微店",
                "url":"http://v.qq.com/"
            }
            ]
        },
        {
            "name":"工具",
            "sub_button":[
            {  
                "type":"view",
                "name":"搜索",
                "url":"http://www.soso.com/"
            },
            {
                "type":"view",
                "name":"视频",
                "url":"http://v.qq.com/"
            },
            {
                "type":"click",
                "name":"赞一下我们",
                "key":"V1001_TODAY_V1001_GOOD"
            }]
        },
        {
            "name":"其他",
            "sub_button":[
            {  
                "type":"click",
                "name":"帮助",
                "key":"V1001_TODAY_HELP"
            },
            {
                "type":"click",
                "name":"关于",
                "key":"V1001_TODAY_ABOUT"
            },
            {
                "type":"view",
                "name":"官方网站",
                "url":"https://yii.li"
            }]
        }
        ]
    }
    `)
)
