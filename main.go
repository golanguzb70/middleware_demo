package main

import "github.com/azizbek/middleware/api"

func main() {
	engine := api.New()
	engine.Run(":8082")
}
