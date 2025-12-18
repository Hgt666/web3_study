package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-api-demo/model"
	"github.com/golang-jwt/jwt/v5"
)

// 配置项
const (
	secretKey     = "test123"
	AccessExpire  = 15 * time.Minute
	RefreshExpire = 7 * 24 * time.Hour
)

type CustomClaims struct {
	jwt.RegisteredClaims        // 继承标准声明
	UserID               uint64 `json:"user_id"`
	UserName             string `json:"username"`
}

// GenerateToken 生成token和RefreshToken
func GenerateToken(UserID uint64, UserName string) (accessToken, refreshToken string, err error) {
	// 1、生成Access Token (短期，包含完整信息)
	accessClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-api-demo",
		},
		UserID:   UserID,
		UserName: UserName,
	}

	// 签名生成Access Token
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", fmt.Errorf("生成Access Token失败: %v", err)
	}

	// 2、生成RefreshToekn（长期，仅保留核心信息）
	refreshClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-api-demo",
		},
		UserID: UserID,
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", fmt.Errorf("生成Refresh Token失败: %v", err)
	}
	return accessToken, refreshToken, nil

}

// ParseToken 解析并验证token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析Token（指定自定义声明和签名验证函数）
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法(防止算法被篡改)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名算法: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil // 返回签名密钥

	})
	if err != nil {
		// 区分Token 过期/无效/解析失败
		return nil, err
	}

	// 验证Token是否有效并提取声明
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的Token或已被篡改")

}

// RefreshAccessToken 刷新AccessToken 通过有效的Refresh Token生成新的AccessToken
func RefreshAccessToken(refreshToken string) (newAccessToken string, err error) {
	// 解析Refresh Token
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("刷新 Token失败: %v", err)
	}
	// 从数据库中获取用户信息
	var user model.User
	if err := DB.First(&user, claims.UserID).Error; err != nil {
		return "", fmt.Errorf("获取用户信息失败: %v", err)
	}
	username := user.UserName

	// 生成新的AccessToken
	newAccessClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-api-demo",
		},
		UserID:   claims.UserID,
		UserName: username,
	}

	newAccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("生成新AccessToken失败: %v", err)
	}
	return newAccessToken, nil

}
