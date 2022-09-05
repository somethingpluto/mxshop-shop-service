package handler

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"order_service/global"
	"order_service/model"
	"order_service/proto"
)

func (s *OrderService) OrderList(ctx context.Context, request *proto.OrderFilterRequest) (*proto.OrderListResponse, error) {
	response := &proto.OrderListResponse{}

	var orders []model.OrderInfo
	var total int64
	global.DB.Where(&model.OrderInfo{User: request.UserId}).Count(&total)
	response.Total = int32(total)
	result := global.DB.Scopes(model.Paginate(int(request.Pages), int(request.PagePerNums))).Find(&orders)
	if result.RowsAffected == 0 {
		zap.S().Warnw("[orderList] 订单分页查询失败")
		return nil, status.Errorf(codes.NotFound, "未查询到订单")
	}
	for _, order := range orders {
		response.Data = append(response.Data, &proto.OrderInfoResponse{
			Id:      order.ID,
			UserId:  order.User,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Status:  order.Status,
			Post:    order.Post,
			Total:   order.OrderMount,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SingerMobile,
		})
	}
	return response, nil
}

// OrderDetail
// @Description: 获取订单详情
// @receiver s
// @param ctx
// @param request
// @return *proto.OrderInfoDetailResponse
// @return error
//
func (s *OrderService) OrderDetail(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoDetailResponse, error) {
	response := &proto.OrderInfoDetailResponse{}
	var order model.OrderInfo

	result := global.DB.Where(&model.OrderInfo{BaseModel: model.BaseModel{ID: request.Id}, User: request.UserId}).First(&order)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	orderInfo := proto.OrderInfoResponse{}
	orderInfo.Id = order.ID
	orderInfo.UserId = order.User
	orderInfo.OrderSn = order.OrderSn
	orderInfo.PayType = order.PayType
	orderInfo.Status = order.Status
	orderInfo.Post = order.Post
	orderInfo.Total = order.OrderMount
	orderInfo.Address = order.Address
	orderInfo.Name = order.SignerName
	orderInfo.Mobile = order.SingerMobile

	response.OrderInfo = &orderInfo

	var orderGoods []model.OrderGoods
	result = global.DB.Where(&model.OrderGoods{Order: order.ID}).Find(&orderGoods)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, orderGood := range orderGoods {
		response.Goods = append(response.Goods, &proto.OrderItemResponse{
			GoodsId:    orderGood.Goods,
			GoodsName:  orderGood.GoodsName,
			GoodsImage: orderGood.GoodsImage,
			Nums:       orderGood.Nums,
		})
	}
	return response, nil
}

func (s *OrderService) CreateOrder(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	response := &proto.OrderInfoResponse{}

	var goodsId []int32
	var shopCarts []model.ShoppingCart
	goodsNumsMap := make(map[int32]int32)
	result := global.DB.Where(&model.ShoppingCart{User: request.UserId, Checked: true}).Find(&shopCarts)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "没有选中的结算商品")
	}

	for _, shopCart := range shopCarts {
		goodsId = append(goodsId, shopCart.Goods)
		goodsNumsMap[shopCart.Goods] = shopCart.Nums
	}
	// 调用商品微服务 查询商品信息
	goods, err := global.GoodsServiceClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: goodsId})
	if err != nil {
		zap.S().Errorw("[goods_service]服务批量查询商品失败", "err", err)
		return nil, status.Errorf(codes.Internal, "批量查询商品信息失败")
	}
	var orderAmount float32
	var orderGoods []*model.OrderGoods
	var goodsInvInfo []*proto.GoodsInvInfo
	for _, goods := range goods.Data {
		orderAmount += goods.ShopPrice * float32(goodsNumsMap[goods.Id])
		orderGoods = append(orderGoods, &model.OrderGoods{
			Goods:      goods.Id,
			GoodsName:  goods.Name,
			GoodsImage: goods.GoodsFrontImage,
			GoodsPrice: goods.ShopPrice,
			Nums:       goodsNumsMap[goods.Id],
		})
		goodsInvInfo = append(goodsInvInfo, &proto.GoodsInvInfo{GoodsId: goods.Id, Num: goodsNumsMap[goods.Id]})
	}
	// 调用库存服务 扣减库存
	_, err = global.InventoryServiceClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: goodsInvInfo,
	})
	if err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "库存服务扣减失败")
	}
	// 生成订单表
	tx := global.DB.Begin()
	order := model.OrderInfo{
		OrderSn:      GenerateOrderSn(request.UserId),
		OrderMount:   orderAmount,
		Address:      request.Address,
		SignerName:   request.Name,
		SingerMobile: request.Mobile,
		Post:         request.Post,
	}
	result = tx.Save(&order)
	if result.Error != nil {
		zap.S().Errorw("保存订单失败", "err", result.Error)
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")

	}
	for _, orderGood := range orderGoods {
		orderGood.Order = order.ID
	}

	// 删除购物车记录
	result = tx.Where(&model.ShoppingCart{User: request.UserId, Checked: true}).Delete(&model.ShoppingCart{})
	if result.RowsAffected == 0 {
		zap.S().Errorw("删除购物车记录失败", "err", result.Error)
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建订单失败")
	}
	tx.Commit()

	response.Id = order.ID
	response.OrderSn = order.OrderSn
	response.Total = order.OrderMount
	return response, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, request *proto.OrderStatus) (*emptypb.Empty, error) {
	result := global.DB.Model(&model.OrderInfo{}).Where(&model.OrderInfo{OrderSn: request.OrderSn}).Update("status", request.Status)
	if result.RowsAffected == 0 || result.Error != nil {
		return nil, status.Errorf(codes.Internal, "更新订单状态失败")
	}
	return &emptypb.Empty{}, nil
}
