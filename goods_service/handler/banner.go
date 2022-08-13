package handler

import (
	"context"
	"goods_service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// BannerList
// @Description: 获取轮播图列表
// @receiver g
// @param ctx
// @return *proto.BannerListResponse
// @return error
//
func (g GoodsServer) BannerList(ctx context.Context, request *emptypb.Empty) (*proto.BannerListResponse, error) {
	response := &proto.BannerListResponse{}
	return response, nil
}

// CreateBanner
// @Description: 创建轮播图
// @receiver g
// @param ctx
// @param request
// @return *proto.BannerResponse
// @return error
//
func (g GoodsServer) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	response := &proto.BannerResponse{}
	return response, nil
}

// DeleteBanner
// @Description: 删除轮播图
// @receiver g
// @param ctx
// @param request
// @return *proto.OperationResult
// @return error
//
func (g GoodsServer) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*proto.OperationResult, error) {
	response := &proto.OperationResult{}
	return response, nil
}

// UpdateBanner
// @Description: 更新轮播图
// @receiver g
// @param ctx
// @param request
// @return *proto.BannerResponse
// @return error
//
func (g GoodsServer) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	response := &proto.BannerResponse{}
	return response, nil
}
