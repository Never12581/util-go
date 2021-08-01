package api

import "github.com/gogf/gf/frame/gmvc"

var DbGeneratorTemplate = &dbGeneratorTemplate{}

type dbGeneratorTemplate struct {
	gmvc.Controller
}

func (c *dbGeneratorTemplate) Index() {
	c.View.Assign("PackageName", "I am your father")
	c.View.Assigns(map[string]interface{}{
		"age":   18,
		"score": 100,
	})
	c.View.Display("html/index.html")
}
