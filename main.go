package main

import (
	_ "bee_hello/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

