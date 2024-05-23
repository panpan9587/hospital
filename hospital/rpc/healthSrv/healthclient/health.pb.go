// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: rpc/doc/health.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 体检信息表
type BodyInspectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	PeopleMsgId int64  `protobuf:"varint,20,opt,name=PeopleMsgId,proto3" json:"PeopleMsgId,omitempty"` //个人信息id
	Height      int64  `protobuf:"varint,30,opt,name=Height,proto3" json:"Height,omitempty"`           //身高
	Weight      int64  `protobuf:"varint,40,opt,name=Weight,proto3" json:"Weight,omitempty"`           //体重
	Inheritance string `protobuf:"bytes,50,opt,name=Inheritance,proto3" json:"Inheritance,omitempty"`  //遗传
	DoctorId    int64  `protobuf:"varint,60,opt,name=DoctorId,proto3" json:"DoctorId,omitempty"`       //医生id
}

func (x *BodyInspectInfo) Reset() {
	*x = BodyInspectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInspectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInspectInfo) ProtoMessage() {}

func (x *BodyInspectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInspectInfo.ProtoReflect.Descriptor instead.
func (*BodyInspectInfo) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{0}
}

func (x *BodyInspectInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BodyInspectInfo) GetPeopleMsgId() int64 {
	if x != nil {
		return x.PeopleMsgId
	}
	return 0
}

func (x *BodyInspectInfo) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BodyInspectInfo) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *BodyInspectInfo) GetInheritance() string {
	if x != nil {
		return x.Inheritance
	}
	return ""
}

func (x *BodyInspectInfo) GetDoctorId() int64 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

// 体检表添加 --入参
type BodyInspectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeopleMsgId int64  `protobuf:"varint,20,opt,name=PeopleMsgId,proto3" json:"PeopleMsgId,omitempty"` //个人信息id
	Height      int64  `protobuf:"varint,30,opt,name=Height,proto3" json:"Height,omitempty"`           //身高
	Weight      int64  `protobuf:"varint,40,opt,name=Weight,proto3" json:"Weight,omitempty"`           //体重
	Inheritance string `protobuf:"bytes,50,opt,name=Inheritance,proto3" json:"Inheritance,omitempty"`  //遗传
	DoctorId    int64  `protobuf:"varint,60,opt,name=DoctorId,proto3" json:"DoctorId,omitempty"`       //医生id
}

func (x *BodyInspectRequest) Reset() {
	*x = BodyInspectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInspectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInspectRequest) ProtoMessage() {}

func (x *BodyInspectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInspectRequest.ProtoReflect.Descriptor instead.
func (*BodyInspectRequest) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{1}
}

func (x *BodyInspectRequest) GetPeopleMsgId() int64 {
	if x != nil {
		return x.PeopleMsgId
	}
	return 0
}

func (x *BodyInspectRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BodyInspectRequest) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *BodyInspectRequest) GetInheritance() string {
	if x != nil {
		return x.Inheritance
	}
	return ""
}

func (x *BodyInspectRequest) GetDoctorId() int64 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

type BodyInspectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BodyInspect *BodyInspectInfo `protobuf:"bytes,10,opt,name=BodyInspect,proto3" json:"BodyInspect,omitempty"`
}

func (x *BodyInspectResponse) Reset() {
	*x = BodyInspectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyInspectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyInspectResponse) ProtoMessage() {}

func (x *BodyInspectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyInspectResponse.ProtoReflect.Descriptor instead.
func (*BodyInspectResponse) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{2}
}

func (x *BodyInspectResponse) GetBodyInspect() *BodyInspectInfo {
	if x != nil {
		return x.BodyInspect
	}
	return nil
}

type MedicalItemsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,10,opt,name=Id,proto3" json:"Id,omitempty"`
	Visual        string `protobuf:"bytes,20,opt,name=Visual,proto3" json:"Visual,omitempty"`               //视力检查
	Hearing       string `protobuf:"bytes,30,opt,name=Hearing,proto3" json:"Hearing,omitempty"`             //听力检查
	BloodPressure string `protobuf:"bytes,40,opt,name=BloodPressure,proto3" json:"BloodPressure,omitempty"` //血压
	BloodSugar    string `protobuf:"bytes,50,opt,name=BloodSugar,proto3" json:"BloodSugar,omitempty"`       //血糖
	Urine         string `protobuf:"bytes,60,opt,name=Urine,proto3" json:"Urine,omitempty"`                 //尿常规
	Ct            int64  `protobuf:"varint,70,opt,name=Ct,proto3" json:"Ct,omitempty"`                      //CT 1.胸部ct 2.脑部ct
}

func (x *MedicalItemsInfo) Reset() {
	*x = MedicalItemsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicalItemsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicalItemsInfo) ProtoMessage() {}

func (x *MedicalItemsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicalItemsInfo.ProtoReflect.Descriptor instead.
func (*MedicalItemsInfo) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{3}
}

func (x *MedicalItemsInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MedicalItemsInfo) GetVisual() string {
	if x != nil {
		return x.Visual
	}
	return ""
}

func (x *MedicalItemsInfo) GetHearing() string {
	if x != nil {
		return x.Hearing
	}
	return ""
}

func (x *MedicalItemsInfo) GetBloodPressure() string {
	if x != nil {
		return x.BloodPressure
	}
	return ""
}

func (x *MedicalItemsInfo) GetBloodSugar() string {
	if x != nil {
		return x.BloodSugar
	}
	return ""
}

func (x *MedicalItemsInfo) GetUrine() string {
	if x != nil {
		return x.Urine
	}
	return ""
}

func (x *MedicalItemsInfo) GetCt() int64 {
	if x != nil {
		return x.Ct
	}
	return 0
}

// 根据id查询体检项目的详情 --入参
type MedicalItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,10,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *MedicalItemsRequest) Reset() {
	*x = MedicalItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicalItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicalItemsRequest) ProtoMessage() {}

func (x *MedicalItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicalItemsRequest.ProtoReflect.Descriptor instead.
func (*MedicalItemsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{4}
}

func (x *MedicalItemsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 反参
type MedicalItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedicalItems *MedicalItemsInfo `protobuf:"bytes,10,opt,name=MedicalItems,proto3" json:"MedicalItems,omitempty"`
}

func (x *MedicalItemsResponse) Reset() {
	*x = MedicalItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicalItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicalItemsResponse) ProtoMessage() {}

func (x *MedicalItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicalItemsResponse.ProtoReflect.Descriptor instead.
func (*MedicalItemsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{5}
}

func (x *MedicalItemsResponse) GetMedicalItems() *MedicalItemsInfo {
	if x != nil {
		return x.MedicalItems
	}
	return nil
}

// 预约体检信息详情 --入参
type GetBodyInspectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,20,opt,name=Id,proto3" json:"Id,omitempty"` //个人信息id
}

func (x *GetBodyInspectRequest) Reset() {
	*x = GetBodyInspectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBodyInspectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBodyInspectRequest) ProtoMessage() {}

func (x *GetBodyInspectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBodyInspectRequest.ProtoReflect.Descriptor instead.
func (*GetBodyInspectRequest) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{6}
}

func (x *GetBodyInspectRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 预约体检信息详情 --反参
type GetBodyInspectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BodyInspect *BodyInspectInfo `protobuf:"bytes,10,opt,name=BodyInspect,proto3" json:"BodyInspect,omitempty"`
}

func (x *GetBodyInspectResponse) Reset() {
	*x = GetBodyInspectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBodyInspectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBodyInspectResponse) ProtoMessage() {}

func (x *GetBodyInspectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBodyInspectResponse.ProtoReflect.Descriptor instead.
func (*GetBodyInspectResponse) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{7}
}

func (x *GetBodyInspectResponse) GetBodyInspect() *BodyInspectInfo {
	if x != nil {
		return x.BodyInspect
	}
	return nil
}

type SignInInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,10,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId       int64  `protobuf:"varint,20,opt,name=UserId,proto3" json:"UserId,omitempty"`            //用户id
	SignInMethod string `protobuf:"bytes,30,opt,name=SignInMethod,proto3" json:"SignInMethod,omitempty"` //签到方式
	Status       int64  `protobuf:"varint,40,opt,name=Status,proto3" json:"Status,omitempty"`            //签到状态
}

func (x *SignInInfo) Reset() {
	*x = SignInInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInInfo) ProtoMessage() {}

func (x *SignInInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInInfo.ProtoReflect.Descriptor instead.
func (*SignInInfo) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{8}
}

func (x *SignInInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SignInInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SignInInfo) GetSignInMethod() string {
	if x != nil {
		return x.SignInMethod
	}
	return ""
}

func (x *SignInInfo) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 签到记录 --入参
type GetSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,20,opt,name=UserId,proto3" json:"UserId,omitempty"`            //用户id
	SignInMethod string `protobuf:"bytes,30,opt,name=SignInMethod,proto3" json:"SignInMethod,omitempty"` //签到方式
	Status       int64  `protobuf:"varint,40,opt,name=Status,proto3" json:"Status,omitempty"`            //签到状态
}

func (x *GetSignInRequest) Reset() {
	*x = GetSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignInRequest) ProtoMessage() {}

func (x *GetSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignInRequest.ProtoReflect.Descriptor instead.
func (*GetSignInRequest) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{9}
}

func (x *GetSignInRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetSignInRequest) GetSignInMethod() string {
	if x != nil {
		return x.SignInMethod
	}
	return ""
}

func (x *GetSignInRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 签到记录 --反参
type GetSignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignIn *SignInInfo `protobuf:"bytes,10,opt,name=SignIn,proto3" json:"SignIn,omitempty"`
}

func (x *GetSignInResponse) Reset() {
	*x = GetSignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_doc_health_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignInResponse) ProtoMessage() {}

func (x *GetSignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_doc_health_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignInResponse.ProtoReflect.Descriptor instead.
func (*GetSignInResponse) Descriptor() ([]byte, []int) {
	return file_rpc_doc_health_proto_rawDescGZIP(), []int{10}
}

func (x *GetSignInResponse) GetSignIn() *SignInInfo {
	if x != nil {
		return x.SignIn
	}
	return nil
}

var File_rpc_doc_health_proto protoreflect.FileDescriptor

var file_rpc_doc_health_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a,
	0x0f, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x3c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0xa4, 0x01, 0x0a, 0x12, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x68,
	0x65, 0x72, 0x69, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x13, 0x42, 0x6f, 0x64, 0x79, 0x49,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x42, 0x6f, 0x64, 0x79,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x22, 0xc0, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x65, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x24,
	0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x50, 0x72, 0x65, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x53, 0x75, 0x67,
	0x61, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x6c, 0x6f, 0x6f, 0x64, 0x53,
	0x75, 0x67, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x72, 0x69, 0x6e, 0x65, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x55, 0x72, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x74,
	0x18, 0x46, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x43, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x4d, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x22, 0x52, 0x0a, 0x14, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x4d, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x51,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x42, 0x6f, 0x64, 0x79,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x22, 0x70, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x66, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x32, 0xb3, 0x02, 0x0a, 0x0d, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x64, 0x79, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x49,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_doc_health_proto_rawDescOnce sync.Once
	file_rpc_doc_health_proto_rawDescData = file_rpc_doc_health_proto_rawDesc
)

func file_rpc_doc_health_proto_rawDescGZIP() []byte {
	file_rpc_doc_health_proto_rawDescOnce.Do(func() {
		file_rpc_doc_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_doc_health_proto_rawDescData)
	})
	return file_rpc_doc_health_proto_rawDescData
}

var file_rpc_doc_health_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rpc_doc_health_proto_goTypes = []interface{}{
	(*BodyInspectInfo)(nil),        // 0: user.BodyInspectInfo
	(*BodyInspectRequest)(nil),     // 1: user.BodyInspectRequest
	(*BodyInspectResponse)(nil),    // 2: user.BodyInspectResponse
	(*MedicalItemsInfo)(nil),       // 3: user.MedicalItemsInfo
	(*MedicalItemsRequest)(nil),    // 4: user.MedicalItemsRequest
	(*MedicalItemsResponse)(nil),   // 5: user.MedicalItemsResponse
	(*GetBodyInspectRequest)(nil),  // 6: user.GetBodyInspectRequest
	(*GetBodyInspectResponse)(nil), // 7: user.GetBodyInspectResponse
	(*SignInInfo)(nil),             // 8: user.SignInInfo
	(*GetSignInRequest)(nil),       // 9: user.GetSignInRequest
	(*GetSignInResponse)(nil),      // 10: user.GetSignInResponse
}
var file_rpc_doc_health_proto_depIdxs = []int32{
	0,  // 0: user.BodyInspectResponse.BodyInspect:type_name -> user.BodyInspectInfo
	3,  // 1: user.MedicalItemsResponse.MedicalItems:type_name -> user.MedicalItemsInfo
	0,  // 2: user.GetBodyInspectResponse.BodyInspect:type_name -> user.BodyInspectInfo
	8,  // 3: user.GetSignInResponse.SignIn:type_name -> user.SignInInfo
	1,  // 4: user.HealthService.AddBodyInspect:input_type -> user.BodyInspectRequest
	4,  // 5: user.HealthService.GetMedicalItems:input_type -> user.MedicalItemsRequest
	6,  // 6: user.HealthService.GetBodyInspect:input_type -> user.GetBodyInspectRequest
	9,  // 7: user.HealthService.GetSignIn:input_type -> user.GetSignInRequest
	2,  // 8: user.HealthService.AddBodyInspect:output_type -> user.BodyInspectResponse
	5,  // 9: user.HealthService.GetMedicalItems:output_type -> user.MedicalItemsResponse
	7,  // 10: user.HealthService.GetBodyInspect:output_type -> user.GetBodyInspectResponse
	10, // 11: user.HealthService.GetSignIn:output_type -> user.GetSignInResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_doc_health_proto_init() }
func file_rpc_doc_health_proto_init() {
	if File_rpc_doc_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_doc_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInspectInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInspectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyInspectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicalItemsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicalItemsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicalItemsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBodyInspectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBodyInspectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignInRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_doc_health_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignInResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_doc_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_doc_health_proto_goTypes,
		DependencyIndexes: file_rpc_doc_health_proto_depIdxs,
		MessageInfos:      file_rpc_doc_health_proto_msgTypes,
	}.Build()
	File_rpc_doc_health_proto = out.File
	file_rpc_doc_health_proto_rawDesc = nil
	file_rpc_doc_health_proto_goTypes = nil
	file_rpc_doc_health_proto_depIdxs = nil
}
