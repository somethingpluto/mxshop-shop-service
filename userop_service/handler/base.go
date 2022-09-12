package handler

import "userop_service/proto"

type UserOpService struct {
	proto.UnimplementedAddressServer
	proto.UnimplementedUserFavoriteServer
	proto.UnimplementedMessageServer
}
