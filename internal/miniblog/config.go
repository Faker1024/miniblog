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
	"github.com/marmotedu/miniblog/internal/pkg/log"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var cfgFile string

// 指定配置文件名
const defaultConfigName = "miniblog"

// 服务配置的默认目录
const recommendedHomeDir = ".miniblog"

func initConfig() {
	if cfgFile != "" {
		//	当配置文件位置被传入，直接使用传入的配置文件
		viper.AddConfigPath(cfgFile)
	} else {
		//	如果未传入配置文件位置，则尝试读取Home文件夹下的配置文件
		home, err := homedir.Dir()
		// 当err不为空，输出错误并退出
		cobra.CheckErr(err)
		//将home加入配置文件搜索路径
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		//将当前目录加入配置文件搜索路径
		viper.AddConfigPath(".")
	}
	//将配置文件格式定义为yaml文件
	viper.SetConfigType("yaml")
	//配置文件名
	viper.SetConfigName(defaultConfigName)
	//读取匹配的环境变量
	viper.AutomaticEnv()
	//设置读取的环境变量前缀，如果为小写则自动转化为大写
	viper.SetEnvPrefix("MINIBLOG")
	//将viper.Get(key)字符串中‘.’和‘-’替换为‘_’
	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	//打印viper当前使用的配置文件，方便后续的Debug.
	fmt.Println("Using config file", viper.ConfigFileUsed())
}

// logOptions 从viper中读取日志配置，构建`*log.Options` 并返回
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
