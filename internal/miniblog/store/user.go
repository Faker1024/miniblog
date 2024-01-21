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
	"context"
	"github.com/marmotedu/miniblog/internal/pkg/model"
	"gorm.io/gorm"
)

// UserStore 定义了user模块在store的实现方法
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
	Get(ctx context.Context, username string) (*model.UserM, error)
	Update(ctx context.Context, user *model.UserM) error
}

// users UserStore的实现
type users struct {
	db *gorm.DB
}

func (u users) Get(ctx context.Context, username string) (*model.UserM, error) {
	var user model.UserM
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u users) Update(ctx context.Context, user *model.UserM) error {
	return u.db.Save(user).Error
}

// 确保users实现了UserStore接口
var _ UserStore = (*users)(nil)

// Create 插入一条user记录
func (u users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

func newUsers(db *gorm.DB) UserStore {
	return &users{db}
}
