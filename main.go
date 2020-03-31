package main

import "github.com/khofesh/simple-go-api/setup"

func main() {
	r := setup.Router()

	r.Run(":8090")
}
