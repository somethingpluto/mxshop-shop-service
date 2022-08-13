package handler

import (
	"context"
	"goods_service/proto"
)

// GoodsList
// @Description: 获取商品列表
// @receiver g
// @param ctx
// @param GoodsListRequest
// @return GoodsListResponse
// @return err
//
func (g GoodsServer) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (response *proto.GoodsListResponse, err error) {
	return &proto.GoodsListResponse{}, err
}

// BatchGetGoods
// @Description:批量获取商品信息
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) BatchGetGoods(ctx context.Context, request *proto.BatchGoodsIdInfo) (response *proto.GoodsListResponse, err error) {
	return &proto.GoodsListResponse{}, err
}

// CreateGoods
// @Description: 创建商品
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) CreateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (response *proto.GoodsInfoResponse, err error) {
	return &proto.GoodsInfoResponse{}, err
}

// DeleteGoods
// @Description: 删除商品
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) DeleteGoods(ctx context.Context, request *proto.DeleteGoodsInfo) (response *proto.OperationResult, err error) {
	return &proto.OperationResult{}, err
}

// UpdateGoods
// @Description: 更新商品信息
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) UpdateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (response *proto.GoodsInfoResponse, err error) {
	return &proto.GoodsInfoResponse{}, err
}

// GetGoodsDetail
// @Description: 获取商品详细信息
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) GetGoodsDetail(ctx context.Context, request *proto.GoodsInfoRequest) (response *proto.GoodsInfoResponse, err error) {
	return &proto.GoodsInfoResponse{}, err
}
