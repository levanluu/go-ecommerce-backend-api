package main

import (
	"example.com/ecommerce/internal/initialize"
)

func main() {
	// r := routers.NewServer()
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	initialize.Run()
}
