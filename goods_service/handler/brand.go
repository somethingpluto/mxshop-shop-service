package handler

import (
	"context"
	"fmt"
	"goods_service/global"
	"goods_service/model"
	"goods_service/proto"
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
	brandListResponse := &proto.BrandListResponse{}
	var brands []model.Brands
	result := global.DB.Find(&brands)
	fmt.Println(result.RowsAffected)
	for _, brand := range brands {
		fmt.Printf("%#v \n", brand)
	}
	return brandListResponse, nil
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
	return response, nil
}
