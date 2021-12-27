package main

import (
	"uriboard/system"
	"uriboard/web"
)

func main() {

	conf := system.LoadConfiguration()
	web.StartServer(conf)
}
