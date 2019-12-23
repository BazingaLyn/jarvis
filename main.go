package main

import (
	_ "github.com/BazingaLyn/jarvis/docs"
	"github.com/BazingaLyn/jarvis/routers"
)

// @title gin实战
// @version 1.0
// @description gin开发实战接口列表

// @contact.name API Support
// @host localhost:8080
// @BasePath
func main() {
	r := routers.Routers()
	r.Run()
}
