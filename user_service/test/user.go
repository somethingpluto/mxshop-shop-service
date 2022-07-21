package main

import (
	"Shop_service/user_service/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func main() {
	Init()
	//TestGetUserList()
	//TestGetUserByMobile()
	//TestGetUserById()
	//TestCreateUser()
	TestUpdateUser()
	conn.Close()
}

func TestGetUserList() {
	response, err := userClient.GetUserList(context.Background(), &proto.PageInfoRequest{PageNum: 0, PageSize: 10})
	if err != nil {
		panic(err)
	}
	for _, user := range response.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		TestCheckPassword(user.Password)
	}
}

func TestCheckPassword(password string) {
	response, err := userClient.CheckPassword(context.Background(), &proto.CheckPasswordRequest{Password: "admin", EncryptedPassword: password})
	if err != nil {
		panic(err)
	}
	fmt.Printf("验证密码%v \n", response.Success)
}

func TestGetUserByMobile() {
	response, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "1234561"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("姓名：%s 手机号：%s 密码：%s", response.NickName, response.Mobile, response.Password)
}

func TestGetUserById() {
	response, err := userClient.GetUserById(context.Background(), &proto.IdRequest{Id: 1})
	if err != nil {
		panic(err)
	}
	fmt.Printf("姓名：%s 手机号：%s 密码：%s", response.NickName, response.Mobile, response.Password)
}

func TestCreateUser() {
	response, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfoRequest{Mobile: "15171618793", NickName: "alex", Password: "1390714"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("姓名：%s 手机号：%s 密码：%s", response.NickName, response.Mobile, response.Password)
}

func TestUpdateUser() {
	response, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfoRequest{
		Id:       11,
		NickName: "login",
		Gender:   "female",
		Birthday: 0,
	})
	if err != nil {
		panic(err)
	}
	if response.Success {
		fmt.Println("更新成功")
	} else {
		fmt.Println("更新失败")
	}
}
