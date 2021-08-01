package main

import (
	_ "util-go/boot"
	_ "util-go/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server("util-go").Run()
}
