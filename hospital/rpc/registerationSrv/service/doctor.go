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
			Pid:        int32(v.Pid),
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

// 排班科室医生列表
func (c *Registration) GetOfficeDoctorListByIdTime(ctx context.Context, in *proto.OfficeDoctorListByIdTimeReq) (*proto.OfficeDoctorListByIdTimeRes, error) {
	var (
		doctorList []*proto.ShiftDoctor
	)
	res, err := mysql.GetOfficeDoctorListByIdTime(int(in.OfficeId), in.ShiftTime)
	if err != nil {
		return &proto.OfficeDoctorListByIdTimeRes{}, err
	}
	for _, v := range res {
		doctorList = append(doctorList, &proto.ShiftDoctor{
			Id:       int32(v.ID),
			OfficeId: int32(v.OfficeID),
			DoctorId: int32(v.DoctorId),
			Data:     v.Date,
			Time:     int32(v.Time),
			Count:    int32(v.Count),
			Status:   int32(v.Status),
		})
	}
	return &proto.OfficeDoctorListByIdTimeRes{
		Data: doctorList,
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
