package handler

import (
	"Shop_service/user_service/global"
	"Shop_service/user_service/model"
	"Shop_service/user_service/proto"
	"Shop_service/user_service/util"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserService struct{}

// GetUserList
// @Description: 获取用户列表
// @receiver s
// @param ctx
// @param pageInfoRequest
// @return *proto.UserListResponse
// @return error
//
// TODO:分页查询还有点问题
func (s *UserService) GetUserList(ctx context.Context, pageInfoRequest *proto.PageInfoRequest) (*proto.UserListResponse, error) {
	// 实例化 response
	response := &proto.UserListResponse{}
	// 获取总行数
	var users []model.User
	result := global.DB.Find(&users)
	response.Total = int32(result.RowsAffected)
	// 分页查询

	pageNum := pageInfoRequest.PageNum
	pageSize := pageInfoRequest.PageSize
	global.DB.Scopes(util.Paginate(int(pageNum), int(pageSize))).Find(&users)
	for _, user := range users {
		userInfoResponse := util.ModelToResponse(user)
		response.Data = append(response.Data, userInfoResponse)
	}
	return response, nil
}

func (s *UserService) GetUserByMobile(ctx context.Context, mobileRequest *proto.MobileRequest) (*proto.UserInfoResponse, error) {
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

func (s *UserService) GetUserById(ctx context.Context, idRequest *proto.IdRequest) (*proto.UserInfoResponse, error) {
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

func (s *UserService) CreateUser(ctx context.Context, createUserInfoRequest *proto.CreateUserInfoRequest) (*proto.UserInfoResponse, error) {
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

func (s UserService) UpdateUser(ctx context.Context, UpdateUserInfoRequest *proto.UpdateUserInfoRequest) (*proto.UpdateResponse, error) {
	response := &proto.UpdateResponse{}
	var user model.User
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
	response.Success = true
	return response, nil
}

func (s UserService) CheckPassword(ctx context.Context, checkPasswordRequest *proto.CheckPasswordRequest) (*proto.CheckPasswordResponse, error) {
	response := &proto.CheckPasswordResponse{}
	password := checkPasswordRequest.Password
	EncryptedPassword := checkPasswordRequest.EncryptedPassword
	response.Success = util.VerifyPassword(EncryptedPassword, password)
	return response, nil
}
