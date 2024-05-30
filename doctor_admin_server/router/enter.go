package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/device"
	"github.com/flipped-aurora/gin-vue-admin/server/router/devicemgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/patientmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/userauth"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Diagnosis   diagnosis.RouterGroup
	Userauth    userauth.RouterGroup
	Patientmgmt patientmgmt.RouterGroup
	Device      device.RouterGroup
	Devicemgmt  devicemgmt.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
