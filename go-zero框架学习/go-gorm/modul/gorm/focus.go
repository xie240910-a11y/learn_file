package gorm

// 定义轮播图结构体
type Focus struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Pic      string `json:"pic"`
	Link     string `json:"link"`
	Position int    `json:"position"`
}

func (Focus) TableName() string {
	return "focus"
}
