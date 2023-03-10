package handler

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"order_service/global"
	"order_service/model"
	"order_service/proto"
	"time"
)

type OrderService struct {
	proto.UnimplementedOrderServer
}

// GenerateOrderSn
// @Description: 生成订单号
// @param userId
// @return string
//
func GenerateOrderSn(userId int32) string {
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(), userId, rand.Intn(90)+10)
	return orderSn
}

// CartItemList
// @Description: 获取用户购物车列表
// @receiver s
// @param ctx
// @param request
// @return *proto.CartItemListResponse
// @return error
//
func (s *OrderService) CartItemList(ctx context.Context, request *proto.UserInfo) (*proto.CartItemListResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CartItemList", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	cartItemListSpan := opentracing.GlobalTracer().StartSpan("CartItemList", opentracing.ChildOf(parentSpan.Context()))
	response := &proto.CartItemListResponse{}
	var shopCarts []model.ShoppingCart
	// 根据UserId 查询购物车
	result := global.DB.Where(&model.ShoppingCart{User: request.Id}).Find(&shopCarts)
	if result.Error != nil {
		zap.S().Errorw("CartItemList failed", "err", result.Error)
		return nil, result.Error
	}
	response.Total = int32(result.RowsAffected)
	for _, shopCart := range shopCarts {
		response.Data = append(response.Data, &proto.ShopCartInfoResponse{
			Id:      shopCart.ID,
			UserId:  shopCart.User,
			GoodsId: shopCart.Goods,
			Nums:    shopCart.Nums,
			Checked: shopCart.Checked,
		})
	}
	cartItemListSpan.Finish()
	return response, nil
}

// CreateCartItem
// @Description: 商品加入购物车
// @receiver s
// @param ctx
// @param request
// @return *proto.ShopCartInfoResponse
// @return error
//
func (s *OrderService) CreateCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.ShopCartInfoResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CreateCartItem", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	createCartItemSpan := opentracing.GlobalTracer().StartSpan("CreateCartItem", opentracing.ChildOf(parentSpan.Context()))
	response := &proto.ShopCartInfoResponse{}
	var shopCart model.ShoppingCart
	result := global.DB.Where(&model.ShoppingCart{Goods: request.GoodsId, User: request.UserId}).First(&shopCart)
	if result.RowsAffected == 1 {
		// 如果记录已经存在
		shopCart.Nums += request.Nums
	} else {
		shopCart.User = request.UserId
		shopCart.Goods = request.GoodsId
		shopCart.Nums = request.Nums
		shopCart.Checked = false
	}
	result = global.DB.Save(&shopCart)
	if result.Error != nil {
		zap.S().Errorw("CreateCartItem save failed", "err", result.Error)
		return nil, result.Error
	}
	response.Id = shopCart.ID
	createCartItemSpan.Finish()
	return response, nil
}

// UpdateCartItem
// @Description: 更新购物车记录 是否选择/数量改变
// @receiver s
// @param ctx
// @param request
// @return *emptypb.Empty
// @return error
//
func (s *OrderService) UpdateCartItem(ctx context.Context, request *proto.CartItemRequest) (*emptypb.Empty, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "UpdateCartItem", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	updateCartItemSpan := opentracing.GlobalTracer().StartSpan("UpdateCartItem", opentracing.ChildOf(parentSpan.Context()))
	var shopCart model.ShoppingCart
	result := global.DB.Where(&model.ShoppingCart{Goods: request.GoodsId, User: request.UserId}).First(&shopCart)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}
	shopCart.Checked = request.Checked
	if request.Nums > 0 {
		shopCart.Nums = request.Nums
	}
	result = global.DB.Save(&shopCart)
	if result.Error != nil {
		zap.S().Errorw("UpdateCartItem save failed", "err", result.Error)
		return nil, result.Error
	}
	updateCartItemSpan.Finish()
	return &emptypb.Empty{}, nil
}

// DeleteCartItem
// @Description: 删除购物车记录
// @receiver s
// @param ctx
// @param request
// @return *emptypb.Empty
// @return error
//
func (s *OrderService) DeleteCartItem(ctx context.Context, request *proto.CartItemRequest) (*emptypb.Empty, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "DeleteCartItem", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	deleteCartItemSpan := opentracing.GlobalTracer().StartSpan("DeleteCartItem", opentracing.ChildOf(parentSpan.Context()))
	var shopCart model.ShoppingCart
	result := global.DB.Where(&model.ShoppingCart{User: request.UserId, Goods: request.GoodsId}).Delete(&shopCart)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}
	deleteCartItemSpan.Finish()
	return &emptypb.Empty{}, nil
}
