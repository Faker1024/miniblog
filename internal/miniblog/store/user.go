package store

import (
	"context"
	"github.com/marmotedu/miniblog/internal/pkg/model"
	"gorm.io/gorm"
)

// users UserStore的实现
type users struct {
	db *gorm.DB
}

// Create 插入一条user记录
func (u users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

// 确保users实现了UserStore接口
var _ UserStore = (*users)(nil)

// UserStore 定义了user模块在store的实现方法
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

func newUsers(db *gorm.DB) UserStore {
	return &users{db}
}
