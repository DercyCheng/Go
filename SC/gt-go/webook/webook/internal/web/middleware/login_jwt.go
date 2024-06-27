package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
	"webook/internal/web"
)

// LoginJWTMiddlewareBuilder 是一个结构体，用于构建 JWT 登录中间件
type LoginJWTMiddlewareBuilder struct {
}

// CheckLogin 是一个中间件函数，用于检查用户是否已经登录
func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		// 对于用户注册和登录的请求，不需要进行登录校验
		if path == "/users/signup" || path == "/users/login" {
			return
		}

		// 从请求头中获取 Authorization 字段
		authCode := ctx.GetHeader("Authorization")
		// 如果 Authorization 字段为空，那么返回 401 Unauthorized
		if authCode == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 将 Authorization 字段按照空格分割，得到两部分：Bearer 和 token
		segs := strings.Split(authCode, " ")
		// 如果分割后的结果不是两部分，那么返回 401 Unauthorized
		if len(segs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 获取 token
		tokenStr := segs[1]
		var uc web.UserClaims
		// 解析 token，得到 UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		// 如果解析 token 出错，或者 token 无效，那么返回 401 Unauthorized
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token == nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 获取 token 的过期时间
		expertTime := uc.ExpiresAt
		// 如果 token 距离过期还有 50 秒，那么就更新 token 的过期时间，并将新的 token 设置到响应头中
		if expertTime.Sub(time.Now()) < time.Second*50 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
			tokenStr, err = token.SignedString(web.JWTKey)
			ctx.Header("x-jwt-token", tokenStr)
			if err != nil {
				log.Println(err)
			}
			// 将 UserClaims 设置到请求上下文中
			ctx.Set("user", uc)
		}
	}
}
