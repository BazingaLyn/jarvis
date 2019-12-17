package main

import (
	_ "github.com/BazingaLyn/jarvis/docs"
	"github.com/BazingaLyn/jarvis/routers"
)

func main() {
	r := routers.Routers()
	r.Run()
}
