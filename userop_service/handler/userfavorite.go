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

// GetFavoriteList
// @Description: 获取用户收藏列表
// @receiver s
// @param ctx
// @param request
// @return *proto.UserFavoriteListResponse
// @return error
//
func (s UserOpService) GetFavoriteList(ctx context.Context, request *proto.UserFavoriteRequest) (*proto.UserFavoriteListResponse, error) {
	zap.S().Infow("Info", "method", "GetFavoriteList", "request", request)
	response := &proto.UserFavoriteListResponse{}

	var userFavoriteList []model.UserFavorite
	result := global.DB.Where(&model.UserFavorite{User: request.UserId, Goods: request.GoodsId}).Find(&userFavoriteList)
	if result.RowsAffected == 0 {
		zap.S().Warnw("Warning", "message", "查询地址数据为空", "request", request.UserId)
	}
	response.Total = int32(result.RowsAffected)

	var responseList []*proto.UserFavoriteResponse
	for _, favorite := range userFavoriteList {
		responseList = append(responseList, &proto.UserFavoriteResponse{
			UserId:  favorite.User,
			GoodsId: favorite.Goods,
		})
	}
	response.Data = responseList
	return response, nil
}

// AddUserFavorite
// @Description: 添加用户收藏
// @receiver s
// @param ctx
// @param request
// @return *emptypb.Empty
// @return error
//
func (s UserOpService) AddUserFavorite(ctx context.Context, request *proto.UserFavoriteRequest) (*emptypb.Empty, error) {
	var userFav model.UserFavorite

	userFav.User = request.UserId
	userFav.Goods = request.GoodsId

	result := global.DB.Save(&userFav)
	if result.Error != nil {
		zap.S().Errorw("Error", "message", "创建地址失败", "err", result.Error)
		return nil, status.Errorf(codes.Internal, "创建地址失败")
	}

	return &emptypb.Empty{}, nil
}

func (s UserOpService) GetUserFavoriteDetail(ctx context.Context, req *proto.UserFavoriteRequest) (*emptypb.Empty, error) {
	var userFavorite model.UserFavorite
	result := global.DB.Where("goods=? and user=?", req.GoodsId, req.UserId).Find(&userFavorite)
	if result.RowsAffected == 0 {
		zap.S().Warnw("Warning", "message", "用户收藏为0")
		return nil, status.Errorf(codes.NotFound, "用户收藏记录不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s UserOpService) DeleteUserFavorite(ctx context.Context, request *proto.UserFavoriteRequest) (*emptypb.Empty, error) {
	result := global.DB.Where("goods = ? and user =?", request.GoodsId, request.UserId).Delete(&model.UserFavorite{})
	if result.Error != nil {
		zap.S().Errorw("Error", "message", "删除用户收藏失败", "err", result.Error)
		return nil, status.Errorf(codes.Internal, "删除用户收藏失败 ")
	}
	return &emptypb.Empty{}, nil
}
