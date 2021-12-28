package main

import (
	"math/rand"
	"time"
	"uriboard/system"
	"uriboard/web"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	system.LoadConfiguration()
	web.StartServer()
}
