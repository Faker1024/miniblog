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

package errno

var (

	// OK 代表请求成功
	OK = &Errno{
		HTTP:    200,
		Code:    "",
		Message: "",
	}
	// InternalServerError 代表所有未知的服务器错误
	InternalServerError = &Errno{
		HTTP:    500,
		Code:    "InternalError",
		Message: "Internal server error",
	}
	// ErrPageNotFound 表示路由不匹配错误
	ErrPageNotFound = &Errno{
		HTTP:    404,
		Code:    "ResourceNotFound.PageNotFound",
		Message: "Page not found",
	}
	// ErrBind 参数绑定错误
	ErrBind = &Errno{
		HTTP:    400,
		Code:    "InvalidParameter.BindError",
		Message: "Error occurred while binding the request body to the struct",
	}
	// ErrInvalidParameter 表示所有验证失败的错误
	ErrInvalidParameter = &Errno{
		HTTP:    400,
		Code:    "InvalidParameter",
		Message: "Parameter verification failed",
	}
)
