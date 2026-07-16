package db

import (
	"context"
	dv1 "mxshop/app/user/srv/data/v1"
	metav1 "mxshop/pkg/common/meta/v1" // 常见公共 struct

	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) *users {
	return &users{
		db: db,
	}
}

func (u *users) List(ctx context.Context, req *metav1.ListMeta) (*dv1.UserDOLIst, error) {
	// u.db.Where()
	return nil, nil
}
