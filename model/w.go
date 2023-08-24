package model

type W struct {
	AppMsgCnt  int `json:"app_msg_cnt"`
	AppMsgList []E `json:"app_msg_list"`
}

type E struct {
	Cover      string `json:"cover"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	UpdateTime int    `json:"update_time"`
}
