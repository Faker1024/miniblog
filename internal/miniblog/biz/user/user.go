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

package user

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/marmotedu/miniblog/internal/miniblog/store"
	"github.com/marmotedu/miniblog/internal/pkg/auth"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	v1 "github.com/marmotedu/miniblog/internal/pkg/miniblog/v1"
	"github.com/marmotedu/miniblog/internal/pkg/model"
	"github.com/marmotedu/miniblog/internal/pkg/token"
	"regexp"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	ChangePassword(ctx context.Context, username string, request *v1.ChangePasswordRequest) error
	Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error)
}

type userBiz struct {
	ds store.IStore
}

func (u userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)
	fmt.Println(userM.Password)
	userM.Password, _ = auth.Encrypt(userM.Password)
	fmt.Println(userM.Password)
	err := u.ds.Users().Create(ctx, &userM)
	if err != nil {
		match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error())
		if match {
			return errno.ErrUserAlreadyExist
		}
	}
	return nil
}

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

// ChangePassword 是UserBiz接口中`ChangePassword`方法的实现
func (b userBiz) ChangePassword(ctx context.Context, username string, request *v1.ChangePasswordRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}
	err = auth.Compare(userM.Password, request.OldPassword)
	if err != nil {
		return errno.ErrPasswordIncorrect
	}
	userM.Password, _ = auth.Encrypt(request.NewPassword)
	err = b.ds.Users().Update(ctx, userM)
	if err != nil {
		return err
	}
	return nil
}

func (b userBiz) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	//获取登录用户全部信息
	user, err := b.ds.Users().Get(ctx, request.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	//对比传入的明文密码和数据库中已加密过的密码是否匹配
	err = auth.Compare(user.Password, request.Password)
	if err != nil {
		return nil, errno.ErrPasswordIncorrect
	}
	//匹配成功，登录成功，，签发token并返回
	t, err := token.Sign(request.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}
	return &v1.LoginResponse{Token: t}, nil
}
