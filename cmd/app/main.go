package main

import "github.com/srjchsv/weatherservice/internal/router"

func main() {
	r := router.RegisterRouter()
	r.Run(":8080")
}
