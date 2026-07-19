package user

import (
	upbv1 "mxshop/api/user/v1"
	srvv1 "mxshop/app/user/srv/service/v1"
)

type userServer struct {
	upbv1.UnimplementedUserServer
	srv srvv1.UserSrv
}

func NewUserServer(srv srvv1.UserSrv) *userServer {
	return &userServer{
		srv: srv,
	}
}

var _ upbv1.UserServer = (*userServer)(nil)
