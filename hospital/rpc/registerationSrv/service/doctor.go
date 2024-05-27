package service

import (
	"context"
	"demo/rpc/model/mysql"
	proto "demo/rpc/registerationSrv/registerationclient"
)

// 科室列表
func (c *Registration) OfficeList(ctx context.Context, in *proto.Empty) (*proto.OfficeListRes, error) {
	var (
		list []*proto.DoctorOffice
	)
	res, err := mysql.GetOfficeList()
	if err != nil {
		return &proto.OfficeListRes{}, err
	}
	for _, v := range res {
		list = append(list, &proto.DoctorOffice{
			Id:         int32(v.ID),
			OfficeName: v.OfficeName,
			Status:     int32(v.Status),
		})
	}
	if err != nil {
		return &proto.OfficeListRes{}, err
	}
	return &proto.OfficeListRes{
		Data: list,
	}, nil
}

// 科室医生列表
func (c *Registration) OfficeDoctorList(ctx context.Context, in *proto.OfficeDoctorListReq) (*proto.OfficeDoctorListRes, error) {
	var (
		doctorList []*proto.Doctor
	)
	res, err := mysql.GetOfficeDoctorList(int(in.OfficeId))
	if err != nil {
		return &proto.OfficeDoctorListRes{}, err
	}
	for _, v := range res.Doctor {
		doctorList = append(doctorList, &proto.Doctor{
			Id:          int32(v.ID),
			DoctorName:  v.DoctorName,
			Age:         int32(v.Age),
			Sex:         int32(v.Sex),
			Position:    v.Position,
			Tag:         v.Tag,
			Description: v.Description,
			WorkAge:     int32(v.WorkAge),
			WorkTime:    v.WorkTime,
			OfficeId:    int32(v.OfficeID),
			Status:      int32(v.Status),
		})
	}
	list := &proto.DoctorOffice{
		Id:         int32(res.ID),
		OfficeName: res.OfficeName,
		Status:     int32(res.Status),
		Doctor:     doctorList,
	}
	return &proto.OfficeDoctorListRes{
		Data: list,
	}, nil
}

// 医生详请
func (c *Registration) DoctorDetails(ctx context.Context, in *proto.DoctorDetailsReq) (*proto.DoctorDetailsRes, error) {
	var ()
	res, err := mysql.GetDoctorById(int(in.DoctorId))
	if err != nil {
		return &proto.DoctorDetailsRes{}, err
	}
	docker := &proto.Doctor{
		Id:          int32(res.ID),
		DoctorName:  res.DoctorName,
		Age:         int32(res.Age),
		Sex:         int32(res.Sex),
		Position:    res.Position,
		Tag:         res.Tag,
		Description: res.Description,
		WorkAge:     int32(res.WorkAge),
		WorkTime:    res.WorkTime,
		OfficeId:    int32(res.OfficeID),
		Status:      int32(res.Status),
	}
	return &proto.DoctorDetailsRes{
		Data: docker,
	}, nil
}
