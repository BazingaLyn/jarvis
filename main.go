package main

import (
	"github.com/BazingaLyn/jarvis/routers"
)

func main() {
	r := routers.Routers()

	//url := ginSwagger.URL("http://localhost:10086/swagger/doc.json") // The url pointing to API definition
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":10086")
}
