package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/health"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Health    health.RouterGroup
	Diagnosis diagnosis.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
