package v1

import (
	"context"
	metav1 "mxshop/pkg/common/meta/v1" // 常见公共 struct
	"time"

	"gorm.io/gorm"
)

// data 层同理自己约束Req和Res，这样也不会对上层形成强耦合。也就是这里并不关心使用grom或者原生sql，上层更不会看到。
// data 层的struct喜欢以 DO 命名
type UserDO struct {
	ID       int32
	Mobile   string
	NickName string
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string
	Role     int32
	Password string
}

type UserDOLIst struct {
	TotalCount int
	UserList   []*UserDO
}

func (UserDO) TableName() string {
	return "users"
}

type DataFactory interface {
	Users() UserStore
	Begin() *gorm.DB
}

type UserStore interface {
	List(ctx context.Context, req *metav1.ListMeta) (*UserDOLIst, error)
}
