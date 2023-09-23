package utils

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	ValidOr string `json:"valid"`
	BaseClaims
	BufferTime int64
	Claims     jwt.RegisteredClaims
}

func (c CustomClaims) Valid() error {
	if c.ValidOr != "valid" {
		return errors.New("not valid")
	}
	return nil
}

type BaseClaims struct {
	Identity string
	Username string
	ID       string
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		ValidOr:    "valid",
		BaseClaims: baseClaims,
		BufferTime: int64(1 * time.Hour / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		Claims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"}, // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间 7天  配置文件
			Issuer:    "blank",                                            // 签名的发行者
		},
	}
	return claims
}

func CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("temp"))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("temp"), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
