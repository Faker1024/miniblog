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

package v1

// CreateUserRequest 制定了`POST/v1/users`接口的请求参数
type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum, required, stringlength(1|255)" `
	Password string `json:"password" valid:"required, stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required, stringlength(1|255)"`
	Email    string `json:"email" valid:"required, email"`
	Phone    string `json:"phone" valid:"required, stringlength(11|11)"`
}

// LoginResponse 制定`POST/login`接口的返回参数
type LoginResponse struct {
	Token string `json:"token"`
}

// ChangePasswordRequest 指定了`POST/v1/users/{name}/change-password` 接口的请求参数
type ChangePasswordRequest struct {
	//旧密码
	OldPassword string `json:"oldPassword" valid:"required, stringlength(6|18)"`
	//新密码
	NewPassword string `json:"newPassword" valid:"required, stringlength(6|18)"`
}

// LoginRequest 指定`POST/login`接口的请求参数
type LoginRequest struct {
	Username string `json:"username" valid:"alphanum, required, stringlength(1|255)"`
	Password string `json:"password" valid:"required, stringlength(6|18)"`
}

// GetUserResponse 指定了 `POST/v1/users` 接口的请求参数
type GetUserResponse UserInfo

// UserInfo 指定用户的详细信息
type UserInfo struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	PostCount int64  `json:"postCount"`
	CreateAt  string `json:"createAt"`
	UpdateAt  string `json:"updateAt"`
}
