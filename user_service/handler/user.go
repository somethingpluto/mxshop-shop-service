package handler

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"time"
	"user_service/global"
	"user_service/model"
	"user_service/proto"
	"user_service/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var serviceName = "【User_Service】"

type UserService struct{}

// GetUserList
// @Description: 获取用户列表
// @receiver s
// @param ctx
// @param pageInfoRequest
// @return *proto.UserListResponse
// @return error
//
func (s *UserService) GetUserList(ctx context.Context, pageInfoRequest *proto.PageInfoRequest) (*proto.UserListResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "GetUserList", "request", pageInfoRequest)
	// 实例化 response
	response := &proto.UserListResponse{}
	// 获取总行数
	var users []model.User
	result := global.DB.Find(&users)
	response.Total = int32(result.RowsAffected)
	// 分页查询
	var pageUsers []model.User
	pageNum := pageInfoRequest.PageNum
	pageSize := pageInfoRequest.PageSize

	offset := util.Paginate(int(pageNum), int(pageSize))

	global.DB.Offset(offset).Limit(int(pageSize)).Find(&pageUsers)

	for _, user := range pageUsers {
		userInfoResponse := util.ModelToResponse(user)
		response.Data = append(response.Data, userInfoResponse)
	}
	fmt.Println("用户列表")
	return response, nil
}

// GetUserByMobile
// @Description: 通过电话号码获取用户信息
// @receiver s
// @param ctx
// @param mobileRequest
// @return *proto.UserInfoResponse
// @return error
//
func (s *UserService) GetUserByMobile(ctx context.Context, mobileRequest *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "GetUserByMobile", "request", mobileRequest)

	response := &proto.UserInfoResponse{}
	var user model.User
	mobile := mobileRequest.Mobile
	result := global.DB.Where("mobile=?", mobile).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "未查找到该用户")
	}

	response = util.ModelToResponse(user)
	return response, nil
}

// GetUserById
// @Description: 通过ID获取用户信息
// @receiver s
// @param ctx
// @param idRequest
// @return *proto.UserInfoResponse
// @return error
//
func (s *UserService) GetUserById(ctx context.Context, idRequest *proto.IdRequest) (*proto.UserInfoResponse, error) {

	zap.S().Infow("Info", "service", serviceName, "method", "GetUserById", "request", idRequest)

	response := &proto.UserInfoResponse{}

	var user model.User
	id := idRequest.Id
	result := global.DB.Where("id=?", id).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "未查找到该用户")
	}
	response = util.ModelToResponse(user)
	return response, nil
}

// CreateUser
// @Description: 创建用户
// @receiver s
// @param ctx
// @param createUserInfoRequest
// @return *proto.UserInfoResponse
// @return error
//
func (s *UserService) CreateUser(ctx context.Context, createUserInfoRequest *proto.CreateUserInfoRequest) (*proto.UserInfoResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CreateUser", "request", createUserInfoRequest)

	response := &proto.UserInfoResponse{}
	mobile := createUserInfoRequest.Mobile

	var user model.User
	result := global.DB.Where("mobile=?", mobile)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Mobile = createUserInfoRequest.Mobile
	user.NickName = createUserInfoRequest.NickName

	password := createUserInfoRequest.Password
	encryptPassword := util.EncryptPassword(password)
	user.Password = encryptPassword

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	response = util.ModelToResponse(user)
	return response, nil
}

// UpdateUser
// @Description: 更新用户
// @receiver s
// @param ctx
// @param UpdateUserInfoRequest
// @return *proto.UpdateResponse
// @return error
//
func (s UserService) UpdateUser(ctx context.Context, UpdateUserInfoRequest *proto.UpdateUserInfoRequest) (*proto.UpdateResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "UpdateUser", "request", UpdateUserInfoRequest)
	parentSpan := opentracing.SpanFromContext(ctx)
	response := &proto.UpdateResponse{}
	var user model.User
	updateUserSpan := opentracing.GlobalTracer().StartSpan("update_user", opentracing.ChildOf(parentSpan.Context()))
	result := global.DB.First(&user, UpdateUserInfoRequest.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	birthDay := time.Unix(int64(UpdateUserInfoRequest.Birthday), 0)
	user.NickName = UpdateUserInfoRequest.NickName
	user.Birthday = &birthDay
	user.Gender = UpdateUserInfoRequest.Gender

	result = global.DB.Save(user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	updateUserSpan.Finish()
	response.Success = true
	return response, nil
}

// CheckPassword
// @Description: 检查用户密码
// @receiver s
// @param ctx
// @param checkPasswordRequest
// @return *proto.CheckPasswordResponse
// @return error
//
func (s UserService) CheckPassword(ctx context.Context, checkPasswordRequest *proto.CheckPasswordRequest) (*proto.CheckPasswordResponse, error) {
	zap.S().Infow("Info", "service", serviceName, "method", "CheckPassword", "request", checkPasswordRequest)

	response := &proto.CheckPasswordResponse{}
	password := checkPasswordRequest.Password
	EncryptedPassword := checkPasswordRequest.EncryptedPassword
	response.Success = util.VerifyPassword(EncryptedPassword, password)
	return response, nil
}
