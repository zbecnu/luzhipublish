package model

type Weixin struct {
	PublishPage string `json:"publish_page"`
}

type First struct {
	TotalCount    int              `json:"total_count"`
	PublishCount  int              `json:"publish_count"`
	MasssendCount int              `json:"masssend_count"`
	PublishList   []PublishListEle `json:"publish_list"`
}
type PublishListEle struct {
	PublishType int    `json:"publish_type"`
	PublishInfo string `json:"publish_info"`
}

type PublisInfoEle struct {
	Appmsgex []AppmsgexEle `json:"appmsgex"`
}

type AppmsgexEle struct {
	Cover      string `json:"cover"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	UpdateTime int    `json:"update_time"`
}
