package main

import (
	_ "chartmuseum-ui/routers"
	"flag"

	"github.com/astaxie/beego"
)

func init() {
	flag.Parse()
}

func main() {
	beego.Run()
}
