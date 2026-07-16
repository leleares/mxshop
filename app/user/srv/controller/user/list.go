package user

import (
	"context"
	"fmt"
	upbv1 "mxshop/api/user/v1"
	srvv1 "mxshop/app/user/srv/service/v1"
	v1 "mxshop/pkg/common/meta/v1"
)

func GetUserList(ctx context.Context, req *upbv1.PageInfo) (*upbv1.UserListResponse, error) {
	resp, err := srvv1.List(context.Background(), &v1.ListMeta{
		Page:     int(req.Pn),
		PageSize: int(req.PSize),
	})
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	var userListResp *upbv1.UserListResponse
	userListResp.Total = int32(resp.TotalCount)

	for _, v := range resp.UserList {
		userListResp.Data = append(userListResp.Data, &upbv1.UserInfoResponse{
			PassWord: v.Name,
		})
	}

	return userListResp, nil
}
