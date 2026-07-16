package v1

import (
	"context"
	"fmt"
	dv1 "mxshop/app/user/srv/data/v1"
	metav1 "mxshop/pkg/common/meta/v1" // 常见公共 struct
)

// service 层一定是自己定义Req和Res，这样才不会对于controller层形成强耦合，就是controller层怎么变用什么框架不重要，都要按照本文件结构来进行调用
// service 层的struct喜欢以 DTO 命名
type UserDTO struct {
	Name string
}

type userService struct {
	userStore dv1.UserStore
}

func (u *userService) NewUserService(us dv1.UserStore) *userService {
	return &userService{
		userStore: us,
	}
}

func (u *userService) List(ctx context.Context, req *metav1.ListMeta) (*dv1.UserDOLIst, error) {

	resp, err := u.List(ctx, req)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}

	var userDTOListResp *UserDTOLIst
	userDTOListResp.TotalCount = resp.TotalCount
	for _, v := range resp.UserList {
		userDTOListResp.UserList = append(userDTOListResp.UserList, &UserDTO{
			Name: v.Name,
		})
	}
	return nil, nil
}

type UserDTOLIst struct {
	TotalCount int
	UserList   []*UserDTO
}
