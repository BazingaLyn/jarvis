package main

import (
	_ "github.com/BazingaLyn/jarvis/docs"
	"github.com/BazingaLyn/jarvis/routers"
)

// @title jarvis swagger
// @version 1.0
// @description golang web项目
// @contact.name bazinga
// @contact.url https://github.com/BazingaLyn
// @contact.email lgl050712@163.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:10086
func main() {
	r := routers.Routers()
	r.Run(":10086")
}
