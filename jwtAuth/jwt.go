package jwtAuth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"myGo/config"
	"time"
)

// CustomClaims 自定义 Payload 信息
type CustomClaims struct {
	Username string // 手机号
	Password string // 用户昵称
	jwt.StandardClaims
}

func NewCustomClaimsDefault(username string, password string) *CustomClaims {
	beforeTime := time.Now().Unix()
	return &CustomClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: beforeTime,                          // 生效时间
			ExpiresAt: beforeTime + config.JwtConf.Timeout, // 失效时间
			Issuer:    "lzscxb",                            // 机构
		},
	}
}

type JWT struct {
	singKey []byte // Jwt 密钥
}

var (
	TokenExpired     = errors.New("Token is expired")        // 令牌过期
	TokenNotValidYet = errors.New("Token not active yet")    // 令牌未生效
	TokenMalformed   = errors.New("that's not even a token") // 令牌不完整
	TokenInvalid     = errors.New("Token is invalid")        // 无效令牌
)

// NewJWT 返回一个JWT 实例
func NewJWT() *JWT {
	return &JWT{
		singKey: []byte(config.JwtConf.SingKey),
	}
}

// CreateToken 创建新的 Token
func (j *JWT) CreateToken(claims CustomClaims) (token string, err error) {
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(j.singKey)
}

// ParseToken 验证 Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	withClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.singKey, nil
	})
	if err != nil {
		// 获取到 Jwt ValidationError 错误类型
		if ve, ok := err.(*jwt.ValidationError); ok {
			//logger.Error("获取到 Jwt ValidationError 原：%v 错误类型:%v", err, ve.Errors)
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { // 令牌不完整
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { // 令牌过期
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { // 令牌还未生效
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, TokenInvalid
	}
	if withClaims == nil {
		return nil, TokenInvalid
	}

	if claims, ok := withClaims.Claims.(*CustomClaims); ok { // 验证成功
		return claims, nil
	}
	return nil, TokenInvalid
}
