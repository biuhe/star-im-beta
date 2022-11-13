package main

import (
	"star-im/src/router"
)

func main() {
	r := router.Router()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run(":8081")
	if err != nil {
		panic("run app error")
	}
}
