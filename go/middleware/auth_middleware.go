package middleware

import (
	"L-cart/database"
	"L-cart/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func OrgAuthMiddleware() gin.HandlerFunc {
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	return func(c *gin.Context) {
		// クッキーからトークンを取得する前にヘッダーをチェック
		tokenString, err := c.Cookie("token")
		if err != nil {
			// Authorizationヘッダーをチェック
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "トークンが見つかりません"})
				c.Abort()
				return
			}
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// typeクレームが"org"でなければorgログイン由来でないトークンとして拒否
		tokenType, _ := claims["type"].(string)
		if tokenType != "org" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		userId, ok := claims["userId"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// ユーザーの存在チェック
		var user models.User
		if err := database.DB.First(&user, uint(userId)).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーが存在しません"})
			c.Abort()
			return
		}

		c.Set("userId", uint(userId))
		c.Next()
	}
}
