package token

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

// Config 包含token包的配置选项
type Config struct {
	key         string
	identityKey string
}

var (
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	config           = Config{
		key:         "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		identityKey: "identityKey",
	}
	once sync.Once
)

// Init 设置包级别的配置config， config会用于本包后面的token签发和解析
func Init(key, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

// Parse 使用指定的密钥key解析token，解析成功返回token上下文，否则报错
func Parse(tokenString, key string) (string, error) {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	//解析失败
	if err != nil {
		return "", err
	}
	var identityKey string
	// 如果解析成功，从token中取出主题
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		identityKey = claims[config.identityKey].(string)
	}
	return identityKey, nil
}

// ParseRequest 从请求头中获取令牌，并将其传递给Parse函数以解析令牌
func ParseRequest(ctx *gin.Context) (string, error) {
	header := ctx.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return "", ErrMissingHeader
	}
	var t string
	//从请求头中取出token
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, config.key)
}

// Sign 使用jwtSecret签发token，token的claims中会存放传入的subject
func Sign(identityKey string) (tokenString string, err error) {
	//Token的内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"nbf":              time.Now().Unix(),
		"iat":              time.Now().Unix(),
		"exp":              time.Now().Add(100000 * time.Hour).Unix(),
	})
	//签发token
	tokenString, err = token.SignedString([]byte(config.key))
	return
}
