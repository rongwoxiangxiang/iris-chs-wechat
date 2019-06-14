package pojo

type Answer struct {
	ErrCode  int      `json:"ret"`
	Msg      string   `json:"msg"`
	DataJson DataJson `json:"data"`
}

type DataJson interface{}
