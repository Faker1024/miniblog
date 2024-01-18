package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	s    *datastore
)

var _ IStore = (*datastore)(nil)

// IStore 定义了Store层需要实现的方法
type IStore interface {
	Users() UserStore
}

// datastore IStore 的具体实现
type datastore struct {
	db *gorm.DB
}

func (d datastore) Users() UserStore {
	return newUsers(d.db)
}

func NewStore(db *gorm.DB) *datastore {
	// 确保 s 只被初始化一次
	once.Do(func() {
		s = &datastore{db}
	})
	return s
}
