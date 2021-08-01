package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var DbGeneratorApi = &dbGeneratorApi{}

type dbGeneratorApi struct{}

func (dbGeneratorApi *dbGeneratorApi) Index(r *ghttp.Request) {
	r.Cookie.Set("theme", "default")
	r.Session.Set("name", "john")
	content := `Config:{{.Config.redis.cache}}, Cookie:{{.Cookie.theme}}, Session:{{.Session.name}}, Query:{{.Query.name}}`
	r.Response.WriteTplContent(content, nil)
}

func (dbGeneratorApi *dbGeneratorApi) PojoTemplate(r *ghttp.Request) {

	r.Response.Writeln("我是你爸爸很伟大")
}
