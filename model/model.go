package model

type Movie struct {
	Id        string
	Name      string
	Score     float32
	Actors    []string
	Directors []string
	MovieTime int16
}

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
