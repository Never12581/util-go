package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"util-go/app/api"
)

func init() {
	s := g.Server("util-go")
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)
		group.ALL("/dbGenerator/api", api.DbGeneratorApi)
	})
	s.BindController("/dbGenerator", api.DbGeneratorTemplate)
}
