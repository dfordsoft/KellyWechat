package models

type ButtonItem struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key,omitempty"`
	Url  string `json:"url,omitempty"`
}

type MenuItem struct {
	Name      string        `json:"name"`
	SubButton []*ButtonItem `json:"sub_button"`
}

type MenuItems struct {
	Button []*MenuItem `json:"button"`
}
