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

package miniblog

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/miniblog/internal/pkg/log"
	"github.com/marmotedu/miniblog/internal/pkg/middleware"
	"github.com/marmotedu/miniblog/internal/pkg/version/verflag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "A go practical project",
		Long:         "",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			//初始化日志
			log.Init(logOptions())
			verflag.PrintAndExitIfRequest()
			defer log.Sync()
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any aruments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default os $HOME/.yaml)")
	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//以下设置使得initConfig函数在每个命令运行时都会被调用已读取配置
	cobra.OnInitialize(initConfig)
	return cmd
}

// 实际业务代码入口
func run() error {
	// 初始化store层
	err := initStore()
	if err != nil {
		return err
	}
	// 设置Gin模式
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), middleware.RequestID(), middleware.Cors, middleware.Secure, middleware.NoCache}
	g.Use(mws...)
	err = installRouters(g)
	if err != nil {
		return err
	}
	//创建HTTP Server 服务器
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	//运行HTTP 服务器
	//打印一条日志， 用来提示HTTP服务已经起来，方便排障
	log.Infow("Start to listening the incoming request on http address", "addr", viper.GetString("addr"))
	go func() {
		err := httpsrv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()
	// 等待终端信号优雅关闭服务器（10秒超时）
	quit := make(chan os.Signal, 1)
	/*
		kill 默认发送 syscall.SIGTERM 信号
		kill -2 发送 syscall.SIGINT 信号 ctrl + c
		kill -9 发送 syscall.SIGKILL 信号 无法捕获
	*/
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) //不会在此处阻塞
	<-quit                                               //阻塞，接收到上述两种信号继续执行
	log.Infow("Shutting down server ...")
	// 创建ctx用于通知服务器 goroutine， 有10秒时间结束当前请求
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	err = httpsrv.Shutdown(ctx)
	if err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	log.Infow("Server exiting")
	return nil
}
