package jwtAuth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"myGo/common/request"
	"myGo/config"
	"time"
)

// CustomClaims 自定义 Payload 信息
type CustomClaims struct {
	Uuid string // 用户昵称
	jwt.StandardClaims
}

func NewCustomClaimsDefault(uuid string) *CustomClaims {
	beforeTime := time.Now().Unix()
	return &CustomClaims{
		Uuid: uuid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: beforeTime, // 生效时间
			Issuer:    "lzscxb",   // 机构
		},
	}
}

type JWT struct {
	singKey []byte // Jwt 密钥
}

// NewJWT 返回一个JWT 实例
func NewJWT() *JWT {
	return &JWT{
		singKey: []byte(config.TokConf.Secret),
	}
}

// CreateToken 创建新的 Token
func (j *JWT) CreateToken(claims CustomClaims) (token string, err error) {
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(j.singKey)
}

// ParseToken 验证 Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	withClaims, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.singKey, nil
	})
	if withClaims != nil {
		claims, ok := withClaims.Claims.(*CustomClaims)
		if ok { // 验证成功
			return claims, nil
		}
	}
	return nil, errors.New("请求访问：" + request.GetRequest().FullPath() + "，认证失败，无法访问系统资源")
}
