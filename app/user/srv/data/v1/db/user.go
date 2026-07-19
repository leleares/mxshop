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
	if req == nil {
		req = &metav1.ListMeta{}
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	var total int64
	if err := u.db.WithContext(ctx).Model(&dv1.UserDO{}).Count(&total).Error; err != nil {
		return nil, err
	}

	var users []*dv1.UserDO
	err := u.db.WithContext(ctx).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return &dv1.UserDOLIst{
		TotalCount: int(total),
		UserList:   users,
	}, nil
}
