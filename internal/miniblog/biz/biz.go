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
