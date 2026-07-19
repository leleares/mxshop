package user

import (
	"context"
	upbv1 "mxshop/api/user/v1"
	v1 "mxshop/pkg/common/meta/v1"
)

func (us *userServer) GetUserList(ctx context.Context, req *upbv1.PageInfo) (*upbv1.UserListResponse, error) {
	resp, err := us.srv.List(ctx, &v1.ListMeta{
		Page:     int(req.Pn),
		PageSize: int(req.PSize),
	})
	if err != nil {
		return nil, err
	}
	userListResp := &upbv1.UserListResponse{
		Total: int32(resp.TotalCount),
	}

	for _, v := range resp.UserList {
		var birthDay uint64
		if v.Birthday != nil && !v.Birthday.IsZero() {
			birthDay = uint64(v.Birthday.Unix())
		}

		userListResp.Data = append(userListResp.Data, &upbv1.UserInfoResponse{
			Id:       v.ID,
			PassWord: v.Password,
			Mobile:   v.Mobile,
			NickName: v.NickName,
			BirthDay: birthDay,
			Gender:   v.Gender,
			Role:     v.Role,
		})
	}

	return userListResp, nil
}
