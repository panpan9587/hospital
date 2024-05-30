package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	//"github.com/flipped-aurora/gin-vue-admin/server/service/health"
	"github.com/flipped-aurora/gin-vue-admin/server/service/doctor"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"

	"github.com/flipped-aurora/gin-vue-admin/server/service/device"
	"github.com/flipped-aurora/gin-vue-admin/server/service/devicemgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/disease"
	"github.com/flipped-aurora/gin-vue-admin/server/service/doctorment"
	"github.com/flipped-aurora/gin-vue-admin/server/service/doctormsg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/doctorrecipe"
	"github.com/flipped-aurora/gin-vue-admin/server/service/patientmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userauth"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	//HealthServiceGroup    health.ServiceGroup
	DiagnosisServiceGroup    diagnosis.ServiceGroup
	DoctorServiceGroup       doctor.ServiceGroup
	DoctormsgServiceGroup    doctormsg.ServiceGroup
	DoctormentServiceGroup   doctorment.ServiceGroup
	DoctorrecipeServiceGroup doctorrecipe.ServiceGroup
	UserauthServiceGroup     userauth.ServiceGroup
	PatientmgmtServiceGroup  patientmgmt.ServiceGroup
	DeviceServiceGroup       device.ServiceGroup
	DevicemgmtServiceGroup   devicemgmt.ServiceGroup
	DiseaseServiceGroup      disease.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
