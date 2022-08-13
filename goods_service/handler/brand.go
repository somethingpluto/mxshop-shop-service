package handler

import (
	"context"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/model"
	"goods_service/proto"
	"goods_service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BrandList
// @Description: 获取品牌列表
// @receiver g
// @param ctx
// @param request
// @return *proto.BrandListResponse
// @return error
//
func (g GoodsServer) BrandList(ctx context.Context, request *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	response := &proto.BrandListResponse{}

	var brands []model.Brands
	result := global.DB.Scopes(util.Paginate(int(request.Pages), int(request.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	response.Total = int32(total)
	var brandList []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponse := proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Name,
		}
		brandList = append(brandList, &brandResponse)
	}
	response.Data = brandList
	return response, nil
}

// CreateBrand
// @Description: 创建品牌
// @receiver g
// @param ctx
// @param request
// @return *proto.BrandInfoResponse
// @return error
//
func (g GoodsServer) CreateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	response := &proto.BrandInfoResponse{}
	result := global.DB.Where("name=?", request.Name).First(&model.Brands{})
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	brand := model.Brands{
		Name: request.Name,
		Logo: request.Logo,
	}
	zap.S().Infof("创建品牌 %#v", brand)
	global.DB.Create(&brand)
	var newBrand model.Brands
	global.DB.Where("name=?", request.Name).First(&newBrand)
	response.Id = newBrand.ID
	response.Name = newBrand.Name
	response.Logo = newBrand.Logo
	return response, nil
}

// DeleteBrand
// @Description: 删除品牌
// @receiver g
// @param ctx
// @param request
// @return *proto.OperationResult
// @return error
//
func (g GoodsServer) DeleteBrand(ctx context.Context, request *proto.BrandRequest) (*proto.OperationResult, error) {
	response := &proto.OperationResult{}
	result := global.DB.Delete(&model.Brands{}, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	response.Success = true
	return response, nil
}

// UpdateBrand
// @Description: 更新品牌信息
// @receiver g
// @param ctx
// @param request
// @return *proto.BrandInfoResponse
// @return error
//
func (g GoodsServer) UpdateBrand(ctx context.Context, request *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	response := &proto.BrandInfoResponse{}

	var brand model.Brands
	result := global.DB.Where("name=?", request.Name).First(&brand)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	if request.Name != "" {
		brand.Name = request.Name
	}
	if request.Logo != "" {
		brand.Logo = request.Logo
	}
	global.DB.Save(brand)
	response.Id = brand.ID
	return response, nil
}
