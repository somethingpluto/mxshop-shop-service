package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"userop_service/global"
	"userop_service/model"
	"userop_service/proto"
)

func (s UserOpService) GetAddressList(ctx context.Context, reqeust *proto.AddressRequest) (*proto.AddressListResponse, error) {
	zap.S().Infow("Info", "method", "GetAddressList", "request", reqeust)
	response := &proto.AddressListResponse{}

	var address []model.Address
	result := global.DB.Where(&model.Address{User: reqeust.UserId}).Find(&address)
	if result.RowsAffected == 0 {
		zap.S().Warnw("Warning", "message", "查询地址数据为空", "request", reqeust.Id)
	}
	response.Total = int32(result.RowsAffected)
	var addressResponse []*proto.AddressResponse
	for _, addre := range address {
		addressResponse = append(addressResponse, &proto.AddressResponse{
			Id:           addre.ID,
			UserId:       addre.User,
			Province:     addre.Province,
			City:         addre.City,
			District:     addre.District,
			Address:      addre.Address,
			SignerName:   addre.SignerName,
			SignerMobile: addre.SignerMobile,
		})
	}
	response.Data = addressResponse
	return response, nil
}

func (s UserOpService) CreateAddress(ctx context.Context, request *proto.AddressRequest) (*proto.AddressResponse, error) {
	zap.S().Infow("Info", "method", "CreateAddress", "request", request)

	var address model.Address
	address.User = request.UserId
	address.Province = request.Province
	address.City = request.City
	address.District = request.District
	address.Address = request.Address
	address.SignerName = request.SignerName
	address.SignerMobile = request.SignerMobile

	result := global.DB.Save(&address)
	if result.Error != nil {
		zap.S().Errorw("Error", "message", "创建地址失败", "err", result.Error)
		return nil, status.Errorf(codes.Internal, "创建地址失败")
	}
	return &proto.AddressResponse{Id: address.ID}, nil
}

func (s UserOpService) DeleteAddress(ctx context.Context, request *proto.AddressRequest) (*emptypb.Empty, error) {
	zap.S().Infow("Info", "method", "DeleteAddress", "request", request)
	result := global.DB.Where("id = ? and user = ?", request.Id, request.UserId).Delete(&model.Address{})
	if result.RowsAffected == 0 {
		zap.S().Warnw("Warning", "message", "查询地址数据为空", "request", request.Id)
		return nil, status.Errorf(codes.NotFound, "收货地址不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s UserOpService) UpdateAddress(ctx context.Context, request *proto.AddressRequest) (*emptypb.Empty, error) {
	zap.S().Infow("Info", "method", "UpdateAddress", "request", request)

	var address model.Address

	if result := global.DB.Where("id=? and user=?", request.Id, request.UserId).First(&address); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}

	if address.Province != "" {
		address.Province = request.Province
	}

	if address.City != "" {
		address.City = request.City
	}

	if address.District != "" {
		address.District = request.District
	}

	if address.Address != "" {
		address.Address = request.Address
	}

	if address.SignerName != "" {
		address.SignerName = request.SignerName
	}

	if address.SignerMobile != "" {
		address.SignerMobile = request.SignerMobile
	}

	result := global.DB.Save(&address)
	if result.Error != nil {
		zap.S().Errorw("Error", "message", "更新地址失败", "err", result.Error)
		return nil, status.Errorf(codes.Internal, "更新地址失败")
	}

	return &emptypb.Empty{}, nil
}
