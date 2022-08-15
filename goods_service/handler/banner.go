package handler

import (
	"context"
	"goods_service/global"
	"goods_service/model"
	"goods_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var banners []model.Banner
	result := global.DB.Find(&banners)
	response.Total = int32(result.RowsAffected)
	var bannerResponse []*proto.BannerResponse
	for _, banner := range banners {
		bannerResponse = append(bannerResponse, &proto.BannerResponse{
			Id:    banner.ID,
			Index: banner.Index,
			Image: banner.Image,
			Url:   banner.Url,
		})
	}
	response.Data = bannerResponse
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

	var banner model.Banner
	banner.Image = request.Image
	banner.Index = request.Index
	banner.Url = request.Url

	global.DB.Create(&banner)

	var insertBanner model.Banner
	result := global.DB.Where(map[string]interface{}{
		"image": request.Image,
		"index": request.Index,
		"url":   request.Url,
	}).First(&insertBanner)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "创建轮播图失败")
	}

	response.Id = insertBanner.ID
	response.Image = insertBanner.Image
	response.Index = insertBanner.Index
	response.Url = insertBanner.Url
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

	result := global.DB.Delete(&model.Banner{}, request.Id)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	response.Success = true
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

	var banner model.Banner
	result := global.DB.First(&banner, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	if request.Url != "" {
		banner.Url = request.Url
	}
	if request.Image != "" {
		banner.Image = request.Image
	}
	if request.Index != 0 {
		banner.Index = request.Index
	}
	global.DB.Save(&banner)

	return response, nil
}
