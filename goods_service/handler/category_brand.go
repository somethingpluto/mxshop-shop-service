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

// CategoryBrandList
// @Description: 获取目录品牌列表
// @receiver g
// @param ctx
// @param request
// @return *proto.CategoryBrandListResponse
// @return error
//
func (g GoodsServer) CategoryBrandList(ctx context.Context, request *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CategoryBrandList", "request", request)

	response := &proto.CategoryBrandListResponse{}

	var categoryBrands []model.GoodsCategoryBrand
	var total int64
	// 获取记录总条数
	global.DB.Find(&model.GoodsCategoryBrand{}).Count(&total)
	response.Total = int32(total)

	// 连表分页查询
	global.DB.Preload("Category").Preload("Brand").Scopes(util.Paginate(int(request.Pages), int(request.PagePerNums))).Find(&categoryBrands)

	var categroyBrandsResponse []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		categroyBrandsResponse = append(categroyBrandsResponse, &proto.CategoryBrandResponse{
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brand.ID,
				Name: categoryBrand.Brand.Name,
				Logo: categoryBrand.Brand.Logo,
			},
		})
	}
	response.Data = categroyBrandsResponse
	return response, nil
}

// GetCategoryBrandList
// @Description: 获取目录下的品牌列表
// @receiver g
// @param ctx
// @param request
// @return *proto.BrandListResponse
// @return error
//
func (g GoodsServer) GetCategoryBrandList(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "GetCategoryBrandList", "request", request)

	response := &proto.BrandListResponse{} // 查询该商品分类是否存在
	var category model.Category
	result := global.DB.Find(&category, request.Id).First(&category)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}
	// 返回该分类下的 所有商品
	var categoryBrands []model.GoodsCategoryBrand
	// TODO:返回数据为空
	result = global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: request.Id}).Find(&categoryBrands)
	if result.RowsAffected > 0 {
		response.Total = int32(result.RowsAffected)
	}

	var brandInfoResponse []*proto.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponse = append(brandInfoResponse, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brand.ID,
			Name: categoryBrand.Brand.Name,
			Logo: categoryBrand.Brand.Logo,
		})
	}
	response.Data = brandInfoResponse
	return response, nil
}

// CreateCategoryBrand
// @Description: 创建目录下的品牌
// @receiver g
// @param ctx
// @param request
// @return *proto.CategoryBrandResponse
// @return error
//
func (g GoodsServer) CreateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CreateCategoryBrand", "request", request)

	response := &proto.CategoryBrandResponse{}

	var category model.Category
	result := global.DB.First(&category, request.CategoryId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brand
	result = global.DB.First(&brand, request.BrandId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand := model.GoodsCategoryBrand{
		CategoryID: request.CategoryId,
		BrandID:    request.BrandId,
	}
	global.DB.Save(&categoryBrand)
	response.Id = categoryBrand.ID
	return response, nil
}

// DeleteCategoryBrand
// @Description: 删除目录下的品牌
// @receiver g
// @param ctx
// @param request
// @return *proto.OperationResult
// @return error
//
func (g GoodsServer) DeleteCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.OperationResult, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "DeleteCategoryBrand", "request", request)

	response := &proto.OperationResult{}

	result := global.DB.Delete(&model.GoodsCategoryBrand{}, request.Id)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	response.Success = true
	return response, nil
}

// UpdateCategoryBrand
// @Description: 更新目录下的品牌信息
// @receiver g
// @param ctx
// @param request
// @return *proto.OperationResult
// @return error
//
func (g GoodsServer) UpdateCategoryBrand(ctx context.Context, request *proto.CategoryBrandRequest) (*proto.OperationResult, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "UpdateCategoryBrand", "request", request)

	response := &proto.OperationResult{}

	result := global.DB.First(&model.GoodsCategoryBrand{}, request.Id)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	result = global.DB.Find(&model.Category{}, request.CategoryId)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.InvalidArgument, "分类不存在")
	}

	result = global.DB.Find(&model.Brand{}, request.BrandId)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	var categoryBrand model.GoodsCategoryBrand
	categoryBrand.CategoryID = request.CategoryId
	categoryBrand.BrandID = request.BrandId
	return response, nil
}
