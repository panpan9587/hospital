package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/appointmentmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/diagnosis"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/doctor"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/doctorment"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/doctormsg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/registermgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"

	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/device"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/devicemgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/disease"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/doctorrecipe"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/healths"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/patientmgmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userauth"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	DiagnosisApiGroup diagnosis.ApiGroup
	//HealthApiGroup    health.ApiGroup

	DoctorApiGroup          doctor.ApiGroup
	DoctormsgApiGroup       doctormsg.ApiGroup
	DoctormentApiGroup      doctorment.ApiGroup
	DoctorrecipeApiGroup    doctorrecipe.ApiGroup
	UserauthApiGroup        userauth.ApiGroup
	PatientmgmtApiGroup     patientmgmt.ApiGroup
	DeviceApiGroup          device.ApiGroup
	DevicemgmtApiGroup      devicemgmt.ApiGroup
	DiseaseApiGroup         disease.ApiGroup
	AppointmentmgmtApiGroup appointmentmgmt.ApiGroup
	RegistermgmtApiGroup    registermgmt.ApiGroup
	HealthsApiGroup         healths.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
