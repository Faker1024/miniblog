// MIT License
//
// Copyright (c) 2024 jack 3361935899@qq.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	S    *datastore
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
		S = &datastore{db}
	})
	return S
}
