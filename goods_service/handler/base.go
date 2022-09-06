package handler

import "goods_service/proto"

// GoodsServer
// @Description: 商品服务结构体
//
type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

var serviceName = "【Goods_Service】"
