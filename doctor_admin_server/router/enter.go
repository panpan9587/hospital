package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/appointmentmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/device"
	"github.com/flipped-aurora/gin-vue-admin/server/router/devicemgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/router/disease"
	"github.com/flipped-aurora/gin-vue-admin/server/router/doctor"
	"github.com/flipped-aurora/gin-vue-admin/server/router/doctorment"
	"github.com/flipped-aurora/gin-vue-admin/server/router/doctormsg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/doctorrecipe"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/healths"
	"github.com/flipped-aurora/gin-vue-admin/server/router/patientmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/registermgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/userauth"
)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	Diagnosis       diagnosis.RouterGroup
	Doctor          doctor.RouterGroup
	Doctormsg       doctormsg.RouterGroup
	Doctorment      doctorment.RouterGroup
	Doctorrecipe    doctorrecipe.RouterGroup
	Userauth        userauth.RouterGroup
	Patientmgmt     patientmgmt.RouterGroup
	Device          device.RouterGroup
	Devicemgmt      devicemgmt.RouterGroup
	Disease         disease.RouterGroup
	Appointmentmgmt appointmentmgmt.RouterGroup
	Registermgmt    registermgmt.RouterGroup
	Healths         healths.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
