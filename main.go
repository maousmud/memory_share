package main

import (
	"memory_share/router"
)

func main() {
	r := router.GetRouter()
	r.Run()
}
