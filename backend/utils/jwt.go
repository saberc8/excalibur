package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 一些常量
var (
	TokenExpired     error = errors.New("token已过期")
	TokenNotValidYet error = errors.New("token未激活")
	TokenMalformed   error = errors.New("token格式错误")
	TokenInvalid     error = errors.New("token无效")
)

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UID int
	jwt.RegisteredClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte `json:"signing_key"`
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(os.Getenv("SECRET")),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// GenerateToken 生成令牌
func GenerateToken(uId int) string {
	j := NewJWT()
	claims := CustomClaims{
		UID: uId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return err.Error()
	}
	return token
}

// RefreshToken 刷新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// NeedRefresh 判断token是否需要刷新(小于12小时时刷新)
func (j *JWT) NeedRefresh(claims *CustomClaims) bool {
	if claims == nil {
		return false
	}
	now := time.Now()
	expireTime := claims.ExpiresAt.Time
	// 剩余时间小于12小时时刷新
	return expireTime.Sub(now).Hours() < 12
}
