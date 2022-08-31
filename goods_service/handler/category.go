package handler

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/model"
	"goods_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	zap.S().Infof("GetAllCategoriesList request:%v \n", request)
	response := &proto.CategoryListResponse{}

	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, err := json.Marshal(&categorys)
	if err != nil {
		zap.S().Errorw("json转换failed", "err", err.Error())
		return nil, err
	}
	response.JsonData = string(b)
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
	zap.S().Infof("GetSubCategory request:%v \n", request)
	response := &proto.SubCategoryListResponse{}

	var category model.Category
	result := global.DB.First(&category, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	response.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}

	var subCategorys []model.Category
	var subCategorysResponse []*proto.CategoryInfoResponse
	global.DB.Where(&model.Category{ParentCategoryID: request.Id}).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategorysResponse = append(subCategorysResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			ParentCategory: subCategory.ParentCategoryID,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
		})
	}
	response.SubCategory = subCategorysResponse
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
	zap.S().Infof("CreateCategory request:%v", request)
	response := &proto.CategoryInfoResponse{}

	cMap := map[string]interface{}{}
	cMap["name"] = request.Name
	cMap["level"] = request.Level
	cMap["is_tab"] = request.IsTab
	if request.Level != 1 {
		cMap["parent_category_id"] = request.ParentCategory
	}
	global.DB.Model(&model.Category{}).Create(cMap)
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
	zap.S().Infof("DeleteCategory request:%v", request)
	response := &proto.OperationResult{}

	result := global.DB.Delete(&model.Category{}, request.Id)
	if result.RowsAffected == 0 {
		response.Success = false
		return response, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	response.Success = true
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
	zap.S().Infof("UpdateCategory request:%v", request)
	response := &proto.CategoryInfoResponse{}
	var category model.Category
	result := global.DB.First(&category, request.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品")
	}
	if request.Name != "" {
		category.Name = request.Name
	}
	if request.Level != 0 {
		category.Level = request.Level
	}
	if request.ParentCategory != 0 {
		category.ParentCategoryID = request.ParentCategory
	}
	if request.IsTab {
		category.IsTab = request.IsTab
	}
	result = global.DB.Save(&category)

	response.Id = category.ID
	response.Name = category.Name
	response.Level = category.Level
	response.IsTab = category.IsTab
	return response, nil
}
