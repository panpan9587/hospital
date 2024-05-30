package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/device"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/devicemgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/patientmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userauth"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	DiagnosisApiGroup   diagnosis.ApiGroup
	UserauthApiGroup    userauth.ApiGroup
	PatientmgmtApiGroup patientmgmt.ApiGroup
	DeviceApiGroup      device.ApiGroup
	DevicemgmtApiGroup  devicemgmt.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
