package main

import (
	"./api"
)

func main() {
	router := api.Router()
	_ = router.Run(":8080")
}