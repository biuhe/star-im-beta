package main

import (
	"star-im/src/main/config"
	"star-im/src/main/router"
)

func main() {
	config.InitConfig()
	config.InitDB()
	config.InitCache()

	r := router.Router()

	// listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
	err := r.Run(":8081")
	if err != nil {
		panic("run app error")
	}
}
