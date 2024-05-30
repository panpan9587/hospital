package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/device"
	"github.com/flipped-aurora/gin-vue-admin/server/service/devicemgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/patientmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userauth"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	DiagnosisServiceGroup   diagnosis.ServiceGroup
	UserauthServiceGroup    userauth.ServiceGroup
	PatientmgmtServiceGroup patientmgmt.ServiceGroup
	DeviceServiceGroup      device.ServiceGroup
	DevicemgmtServiceGroup  devicemgmt.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
