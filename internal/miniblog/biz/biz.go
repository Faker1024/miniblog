package biz

import (
	"github.com/marmotedu/miniblog/internal/miniblog/biz/user"
	"github.com/marmotedu/miniblog/internal/miniblog/store"
)

// 确保biz实现IBiz接口
var _ IBiz = (*biz)(nil)

// IBiz 定义了Biz层要实现的方法
type IBiz interface {
	User() user.UserBiz
}

// biz 实现IBiz接口
type biz struct {
	ds store.IStore
}

// User 返回一个实现了UserBiz接口的实例
func (b biz) User() user.UserBiz {
	return user.New(b.ds)
}

// NewBiz 创建一个IBiz类型的实例
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}
