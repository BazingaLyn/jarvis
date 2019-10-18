package model

type Movie struct {
	Id        string
	Name      string
	Score     float32
	Actors    []string
	Directors []string
	MovieTime int16
}

type ElasticMovie struct {
	Id         int
	MovieName  string
	Score      float64
	Type       []string
	Directors  []string
	Actors     []string
	Nations    []string
	Languages  []string
	FileLength int
	Describe   string
}

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
