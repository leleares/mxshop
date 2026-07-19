package v1

import (
	"context"
	dv1 "mxshop/app/user/srv/data/v1"
	metav1 "mxshop/pkg/common/meta/v1" // 常见公共 struct
	"time"
)

type UserSrv interface {
	List(ctx context.Context, req *metav1.ListMeta) (*UserDTOLIst, error)
}

// service 层一定是自己定义Req和Res，这样才不会对于controller层形成强耦合，就是controller层怎么变用什么框架不重要，都要按照本文件结构来进行调用
// service 层的struct喜欢以 DTO 命名
type UserDTO struct {
	ID       int32
	Mobile   string
	NickName string
	Birthday *time.Time
	Gender   string
	Role     int32
	Password string
}

type UserDTOLIst struct {
	TotalCount int
	UserList   []*UserDTO
}

type userService struct {
	data dv1.DataFactory
}

func NewUserService(data dv1.DataFactory) UserSrv {
	return &userService{
		data: data,
	}
}

func (u *userService) List(ctx context.Context, req *metav1.ListMeta) (*UserDTOLIst, error) {
	resp, err := u.data.Users().List(ctx, req)
	if err != nil {
		return nil, err
	}

	userDTOListResp := &UserDTOLIst{
		TotalCount: resp.TotalCount,
	}
	for _, v := range resp.UserList {
		userDTOListResp.UserList = append(userDTOListResp.UserList, &UserDTO{
			ID:       v.ID,
			Mobile:   v.Mobile,
			NickName: v.NickName,
			Birthday: v.Birthday,
			Gender:   v.Gender,
			Role:     v.Role,
			Password: v.Password,
		})
	}
	return userDTOListResp, nil
}

var _ UserSrv = (*userService)(nil)
