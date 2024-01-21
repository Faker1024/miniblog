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

package log

import (
	"context"
	"fmt"
	"github.com/marmotedu/miniblog/internal/pkg/know"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

type zapLogger struct {
	z *zap.Logger
}

// 确保 zapLogger 实现了 Logger 接口. 以下变量赋值，可以使错误在编译期被发现.
var _ Logger = &zapLogger{}

var (
	mu sync.Mutex
	//	std定义了默认的全局Logger
	std = NewLogger(NewOptions())
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = NewLogger(opts)
}

func NewLogger(opts *Options) *zapLogger {
	if opts != nil {
		opts = NewOptions()
	}
	//将文本格式的日志级别，例如info转化成对应的 zapcore.Level 类型以供后面使用
	var zapLevel zapcore.Level
	err := zapLevel.UnmarshalText([]byte(opts.Level))
	if err != nil {
		// 如果指定的非法level，则默认使用info级别
		zapLevel = zapcore.InfoLevel
	}
	//创建一个默认的encode配置
	encoderConfig := zap.NewProductionEncoderConfig()
	//自定义MessageKey为message
	encoderConfig.MessageKey = "message"
	//自定义TimeKey为timestamp
	encoderConfig.TimeKey = "timestamp"
	//指定序列化函数，指定序列化时间格式
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	//指定time.Duration序列化函数，将time.Duration序列化经过的毫秒数的浮点数
	//毫秒比默认的秒数更加准确
	encoderConfig.EncodeDuration = func(d time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendFloat64(float64(d) / float64(time.Millisecond))
	}
	// 构建zap.Logger需要的配置
	cfg := &zap.Config{
		// 指定日志级别
		Level: zap.NewAtomicLevelAt(zapLevel),
		// 是否禁止在panic及以上级别打印堆栈信息
		DisableCaller: opts.DisableCaller,
		//是否禁止在panic及以上时打印堆栈信息
		DisableStacktrace: opts.DisableStacktrace,
		//指定日志格式
		Encoding:      opts.Format,
		EncoderConfig: encoderConfig,
		//指定日志输出格式
		OutputPaths: opts.OutputPaths,
		//设置zap内部错误输出位置
		ErrorOutputPaths: []string{"stderr"},
	}
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{z: z}
	// 把标准库的log.Logger的info级别重定向到zap.Logger
	zap.RedirectStdLog(z)
	return logger
}

// C 解析传入的context，尝试提取关注的键值，并添加到zap.Logger结构化日志中
func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

// Sync 调用底层Sync方法，将缓存中的日志刷新入磁盘文件，主程序需要在退出前调用Sync
func Sync() {
	std.Sync()
}

func (l *zapLogger) Sync() {
	err := l.z.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

// Debugw 输出debug级别的日志
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

// Infow 输出debug级别的日志
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

// Warnw 输出debug级别的日志
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw 输出debug级别的日志
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

// Panicw 输出debug级别的日志
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}

// Fatalw 输出debug级别的日志
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	lc := l.clone()
	//将请求Id加入打印日志
	requestID := ctx.Value(know.XRequestIDKey)
	if requestID != nil {
		lc.z = lc.z.With(zap.Any(know.XRequestIDKey, requestID))
	}
	//将用户Id加入打印日志
	userId := ctx.Value(know.XUsernameKey)
	if userId != nil {
		lc.z = lc.z.With(zap.Any(know.XUsernameKey, userId))
	}
	return lc
}

// clone 深度拷贝 zapLogger
func (l *zapLogger) clone() *zapLogger {
	lc := *l
	return &lc
}
