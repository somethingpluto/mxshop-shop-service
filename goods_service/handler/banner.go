package handler

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
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
func (g *GoodsServer) BannerList(ctx context.Context, request *emptypb.Empty) (*proto.BannerListResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "BannerList", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	bannerListSpan := opentracing.GlobalTracer().StartSpan("BannerList", opentracing.ChildOf(parentSpan.Context()))
	response := &proto.BannerListResponse{}
	var banners []model.Banner
	result := global.DB.Find(&banners)
	response.Total = int32(result.RowsAffected)
	bannerListSpan.Finish()
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
func (g *GoodsServer) CreateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CreateBanner", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	createBannerSpan := opentracing.GlobalTracer().StartSpan("CreateBanner", opentracing.ChildOf(parentSpan.Context()))
	response := &proto.BannerResponse{}

	var banner model.Banner
	banner.Image = request.Image
	banner.Index = request.Index
	banner.Url = request.Url

	result := global.DB.Create(&banner)
	if result.Error != nil {
		zap.S().Errorw("创建轮播图失败", "err", result.Error.Error())
		return nil, result.Error
	}
	createBannerSpan.Finish()
	response.Id = banner.ID
	response.Url = banner.Url
	response.Image = banner.Image
	response.Index = banner.Index
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
func (g *GoodsServer) DeleteBanner(ctx context.Context, request *proto.BannerRequest) (*proto.OperationResult, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "DeleteBanner", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	deleteBannerSpan := opentracing.GlobalTracer().StartSpan("DeleteBanner", opentracing.ChildOf(parentSpan.Context()))
	response := &proto.OperationResult{}

	result := global.DB.Delete(&model.Banner{}, request.Id)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	deleteBannerSpan.Finish()
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
func (g *GoodsServer) UpdateBanner(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "UpdateBanner", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	updateBannerSpan := opentracing.GlobalTracer().StartSpan("UpdateBanner", opentracing.ChildOf(parentSpan.Context()))
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
	result = global.DB.Save(&banner)
	if result.RowsAffected != 1 {
		return nil, result.Error
	}
	updateBannerSpan.Finish()
	response.Id = banner.ID
	response.Index = banner.Index
	response.Url = banner.Url
	response.Image = banner.Image
	return response, nil
}
