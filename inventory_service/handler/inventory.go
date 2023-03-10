package handler

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"inventory_service/global"
	"inventory_service/model"
	"inventory_service/proto"
	"time"
)

var serviceName = "【Inventory_Service】"

type InventoryService struct {
	proto.UnimplementedInventoryServer
}

func (i *InventoryService) SetInv(ctx context.Context, request *proto.GoodsInvInfo) (*empty.Empty, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "SetInv", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	setInventorySpan := opentracing.GlobalTracer().StartSpan("setInventory", opentracing.ChildOf(parentSpan.Context()))
	var inventory model.Inventory
	global.DB.Where(&model.Inventory{Goods: request.GoodsId}).First(&inventory)
	inventory.Goods = request.GoodsId
	inventory.Stocks = request.Num
	global.DB.Save(&inventory)
	setInventorySpan.Finish()
	return &empty.Empty{}, nil
}

func (i InventoryService) InvDetail(ctx context.Context, request *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "InvDetail", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	inventoryDetailSpan := opentracing.GlobalTracer().StartSpan("inventoryDetailSpan", opentracing.ChildOf(parentSpan.Context()))
	response := &proto.GoodsInvInfo{}

	var inventory model.Inventory
	result := global.DB.Where(&model.Inventory{
		Goods: request.GoodsId,
	}).First(&inventory)
	if result.RowsAffected == 0 {
		zap.S().Errorw("global.DB.First result = 0", "err", "商品不存在")
		return nil, status.Errorf(codes.NotFound, "库存信息不存在")
	}
	if result.Error != nil {
		zap.S().Errorw("global.DB.First result error", "err", result.Error)
		return nil, status.Errorf(codes.Internal, "数据库查询错误")
	}
	inventoryDetailSpan.Finish()
	response.Num = inventory.Stocks
	response.GoodsId = inventory.Goods
	return response, nil
}

func (i *InventoryService) Sell(ctx context.Context, request *proto.SellInfo) (*empty.Empty, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "Sell", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	sellSpan := opentracing.GlobalTracer().StartSpan("sell", opentracing.ChildOf(parentSpan.Context()))
	tx := global.DB.Begin()
	for _, goodInfo := range request.GoodsInfo {
		mutex := global.Redsync.NewMutex(fmt.Sprintf("goods_%d", goodInfo.GoodsId), redsync.WithTries(100), redsync.WithExpiry(time.Second*20))
		err := mutex.Lock()
		if err != nil {
			zap.S().Errorw("redisync锁错误", "goods_id", goodInfo.GoodsId, "err", err)
			return nil, status.Errorf(codes.Internal, "内部错误")
		}

		var inventory model.Inventory
		result := global.DB.Where(&model.Inventory{
			Goods: goodInfo.GoodsId,
		}).First(&inventory)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品库存信息不存在")
		}
		if inventory.Stocks < goodInfo.Num {
			tx.Rollback() // 回滚之前的操作
			return nil, status.Errorf(codes.ResourceExhausted, "商品库存不足")
		}
		// 扣减
		inventory.Stocks -= goodInfo.Num
		result = tx.Save(&inventory)
		if result.Error != nil {
			return nil, status.Errorf(codes.Internal, "内部错误")
		}
		ok, err := mutex.Unlock()
		if !ok || err != nil {
			zap.S().Errorw("redisync解锁失败", "goods_id", goodInfo.GoodsId, "err", err.Error())
			return nil, status.Errorf(codes.Internal, "内部错误")
		}
	}
	tx.Commit()
	sellSpan.Finish()
	return &empty.Empty{}, nil
}

func (i *InventoryService) ReBack(ctx context.Context, request *proto.SellInfo) (*empty.Empty, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "ReBack", "request", request)
	parentSpan := opentracing.SpanFromContext(ctx)
	rebackSpan := opentracing.GlobalTracer().StartSpan("reback", opentracing.ChildOf(parentSpan.Context()))
	// 库存归还
	// 1.订单超时归还
	// 2.订单创建失败 归还之前扣减的归还
	// 3.手动归还
	tx := global.DB
	for _, goodsInvInfo := range request.GoodsInfo {
		var inventory model.Inventory
		result := global.DB.Where(&model.Inventory{
			Goods: goodsInvInfo.GoodsId,
		}).First(&inventory)
		if result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.NotFound, "商品库存信息不存在")
		}
		inventory.Stocks += goodsInvInfo.Num
		tx.Save(&inventory)
	}
	tx.Commit()
	rebackSpan.Finish()
	return &empty.Empty{}, nil
}
