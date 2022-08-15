package handler

import (
	"context"
	"goods_service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetAllCategoriesList
// @Description: 获取目录列表
// @receiver g
// @param ctx
// @return response
// @return err
//
func (g GoodsServer) GetAllCategoriesList(ctx context.Context, request *emptypb.Empty) (*proto.CategoryListResponse, error) {
	response := &proto.CategoryListResponse{}

	return response, nil
}

// GetSubCategory
// @Description: 获取二级目录
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) GetSubCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	response := &proto.SubCategoryListResponse{}
	return response, nil
}

// CreateCategory
// @Description: 创建目录
// @receiver g
// @param ctx
// @param request
// @return *proto.CategoryInfoResponse
// @return error
//
func (g GoodsServer) CreateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	response := &proto.CategoryInfoResponse{}
	return response, nil
}

// DeleteCategory
// @Description: 删除目录
// @receiver g
// @param ctx
// @param request
// @return *proto.OperationResult
// @return error
//
func (g GoodsServer) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*proto.OperationResult, error) {
	response := &proto.OperationResult{}
	return response, nil
}

// UpdateCategory
// @Description: 更新目录信息
// @receiver g
// @param ctx
// @param request
// @return *proto.CategoryInfoResponse
// @return error
//
func (g GoodsServer) UpdateCategory(ctx context.Context, request *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	response := &proto.CategoryInfoResponse{}
	return response, nil
}
