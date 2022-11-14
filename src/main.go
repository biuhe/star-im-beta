package main

import (
	"star-im/src/config"
	"star-im/src/router"
)

func main() {
	config.InitConfig()
	config.InitDB()

	r := router.Router()

	// listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
	err := r.Run(":8081")
	if err != nil {
		panic("run app error")
	}
}
