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
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/marmotedu/miniblog/internal/pkg/log"
	"github.com/marmotedu/miniblog/internal/pkg/version/verflag"
	__ "github.com/marmotedu/miniblog/proto"
	"github.com/spf13/cobra"
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
	//// 设置Gin模式
	//gin.SetMode(viper.GetString("runmode"))
	//g := gin.New()
	////404页面
	//g.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"code": 10003, "message": "Page not found"})
	//})
	////注册/healthz handler
	//g.GET("/healthz", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	//})
	////创建HTTP Server 服务器
	//httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	//
	////运行HTTP 服务器
	////打印一条日志， 用来提示HTTP服务已经起来，方便排障
	//log.Infow("Start to listening the incoming request on http address", "addr", viper.GetString("addr"))
	//err := httpsrv.ListenAndServe()
	//if err != nil && !errors.Is(err, http.ErrServerClosed) {
	//	log.Fatalw(err.Error())
	//}

	request := &__.HelloWorldRequest{Referer: "hello world"}
	marshal, err := proto.Marshal(request)
	if err != nil {
		log.Fatalw("序列化失败， %d", err)
	}
	i := &__.HelloWorldRequest{}
	err = proto.Unmarshal(marshal, i)
	if err != nil {
		log.Fatalw("反序列化失败， %d", err)
	}
	fmt.Println(i.GetReferer())

	return nil
}
