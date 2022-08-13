package handler

import (
	"context"
	"goods_service/proto"
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
	response := &proto.CategoryBrandListResponse{}
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
	response := &proto.BrandListResponse{}
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
	response := &proto.CategoryBrandResponse{}
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
	response := &proto.OperationResult{}
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
	response := &proto.OperationResult{}
	return response, nil
}
