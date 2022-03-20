package httpServer

type RespStruct struct {
	Status int `json:"status"`
	Data   struct {
		Page `json:"page"`
		Data interface{} `json:"data"`
	} `json:"data"`
	ErrMsg string
}
type Page struct {
	Page       uint `json:"page"`
	Count      uint `json:"count"`
	TotalCount uint `json:"totalCount"`
}
