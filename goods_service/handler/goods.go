package handler

import (
	"context"
	"fmt"
	"goods_service/global"
	"goods_service/model"
	"goods_service/proto"
	"goods_service/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ModelToResponse
// @Description: model数据结构 转换成response数据结构
// @param goods
// @return proto.GoodsInfoResponse
//
func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.Brand.ID,
			Name: goods.Brand.Name,
			Logo: goods.Brand.Logo,
		},
	}
}

// GoodsList
// @Description: 获取商品列表
// @receiver g
// @param ctx
// @param GoodsListRequest
// @return GoodsListResponse
// @return err
//
func (g GoodsServer) GoodsList(ctx context.Context, request *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	response := &proto.GoodsListResponse{}
	zap.S().Infof("GoodsList request: %v", request)
	var goodsList []model.Goods
	localDB := global.DB.Model(&model.Goods{})
	if request.KeyWords != "" {
		localDB = localDB.Where("name LIKE ?", "%"+request.KeyWords+"%")
	}
	if request.IsHot {
		localDB = localDB.Where("is_hot=true")
	}
	if request.IsNew {
		localDB = localDB.Where("is_new=true")
	}
	if request.PriceMin > 0 {
		localDB = localDB.Where("shop_price>=?", request.PriceMin)
	}
	if request.PriceMax > 0 {
		localDB = localDB.Where("shop_price<=?", request.PriceMax)
	}
	if request.Brand > 0 {
		localDB = localDB.Where("brand_id=?", request.Brand)
	}
	// 通过category查询
	var subQuery string
	if request.TopCategory > 0 {
		var category model.Category
		result := global.DB.First(&category, request.TopCategory)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", request.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", request.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", request.TopCategory)
		}
		localDB = localDB.Where(fmt.Sprintf("category_id in (%s)", subQuery))
	}
	var count int64
	localDB.Count(&count)
	response.Total = int32(count)

	result := localDB.Scopes(util.Paginate(int(request.Pages), int(request.PagePerNums))).Find(&goodsList)
	if result.Error != nil {
		return nil, result.Error
	}

	var goodsListResponse []*proto.GoodsInfoResponse
	for _, goods := range goodsList {
		goodsResponse := ModelToResponse(goods)
		goodsListResponse = append(goodsListResponse, &goodsResponse)
	}
	response.Data = goodsListResponse
	return response, nil
}

// BatchGetGoods
// @Description:批量获取商品信息
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) BatchGetGoods(ctx context.Context, request *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	response := &proto.GoodsListResponse{}
	zap.S().Infof("BatchGetGoods request:%v", request)
	var goodsList []model.Goods
	result := global.DB.Where(request.Id).Find(&goodsList)
	var goodsListResponse []*proto.GoodsInfoResponse
	for _, goods := range goodsList {
		goodsInfoResponse := ModelToResponse(goods)
		goodsListResponse = append(goodsListResponse, &goodsInfoResponse)
	}
	response.Total = int32(result.RowsAffected)
	response.Data = goodsListResponse
	return response, nil
}

// CreateGoods
// @Description: 创建商品
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) CreateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	zap.S().Infof("CreateGoods request:%v", request)
	var category model.Category
	if result := global.DB.First(&category, request.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brand
	if result := global.DB.First(&brand, request.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	//先检查redis中是否有这个token
	//防止同一个token的数据重复插入到数据库中，如果redis中没有这个token则放入redis
	//这里没有看到图片文件是如何上传， 在微服务中 普通的文件上传已经不再使用
	goods := model.Goods{
		Brand:           brand,
		BrandID:         brand.ID,
		Category:        category,
		CategoryID:      category.ID,
		Name:            request.Name,
		GoodsSn:         request.GoodsSn,
		MarketPrice:     request.MarketPrice,
		ShopPrice:       request.ShopPrice,
		GoodsBrief:      request.GoodsBrief,
		ShipFree:        request.ShipFree,
		Images:          request.Images,
		DescImages:      request.DescImages,
		GoodsFrontImage: request.GoodsFrontImage,
		IsNew:           request.IsNew,
		IsHot:           request.IsHot,
		OnSale:          request.OnSale,
		Stocks:          request.Stocks,
	}
	global.DB.Create(&goods)
	response := ModelToResponse(goods)
	return &response, nil
}

// DeleteGoods
// @Description: 删除商品
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) DeleteGoods(ctx context.Context, request *proto.DeleteGoodsInfo) (*proto.OperationResult, error) {
	response := &proto.OperationResult{}
	zap.S().Infof("DeleteGoods request:%v", request)
	result := global.DB.Delete(&model.Goods{BaseModel: model.BaseModel{ID: request.Id}}, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	response.Success = true
	return response, nil
}

// UpdateGoods
// @Description: 更新商品信息
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) UpdateGoods(ctx context.Context, request *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error) {
	var goods model.Goods

	if result := global.DB.First(&goods, request.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var category model.Category
	if result := global.DB.First(&category, request.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brand
	if result := global.DB.First(&brand, request.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	goods.Brand = brand
	goods.BrandID = brand.ID
	goods.Category = category
	goods.CategoryID = category.ID
	goods.Name = request.Name
	goods.GoodsSn = request.GoodsSn
	goods.MarketPrice = request.MarketPrice
	goods.ShopPrice = request.ShopPrice
	goods.GoodsBrief = request.GoodsBrief
	goods.ShipFree = request.ShipFree
	goods.Images = request.Images
	goods.DescImages = request.DescImages
	goods.GoodsFrontImage = request.GoodsFrontImage
	goods.IsNew = request.IsNew
	goods.IsHot = request.IsHot
	goods.OnSale = request.OnSale
	global.DB.Save(&goods)
	response := ModelToResponse(goods)
	return &response, nil
}

// GetGoodsDetail
// @Description: 获取商品详细信息
// @receiver g
// @param ctx
// @param request
// @return response
// @return err
//
func (g GoodsServer) GetGoodsDetail(ctx context.Context, request *proto.GoodsInfoRequest) (*proto.GoodsInfoResponse, error) {
	zap.S().Infof("GetGoodsDetail request:%v", request)
	var goods model.Goods
	result := global.DB.Preload("Category").Preload("Brands").First(&goods, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	response := ModelToResponse(goods)
	return &response, nil
}
