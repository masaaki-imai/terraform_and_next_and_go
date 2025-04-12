package org

import (
	"L-cart/database"
	"L-cart/models"
	"L-cart/utils/validation"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleOrgLogin(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if validation.HandleValidationErrors(c, err) {
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"isValid": false,
			"result":  []string{"リクエストの解析に失敗しました"},
		})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isValid": false,
			"result":  []string{"メールアドレスまたはパスワードが正しくありません"},
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isValid": false,
			"result":  []string{"メールアドレスまたはパスワードが正しくありません"},
		})
		return
	}

	// 最新のログイン時間に更新
	now := time.Now()
	user.LastTimeLogin = &now
	database.DB.Save(&user)

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"type":   "org",
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークン生成に失敗しました"})
		return
	}

	c.Header("Set-Cookie", fmt.Sprintf("token=%s; Path=/; Max-Age=%d;", tokenString, 3600*24))
	c.JSON(http.StatusOK, gin.H{
		"isValid": true,
		"result":  true,
	})
}

func HandleLogout(c *gin.Context) {
	// クッキーのトークンをクリア
	c.Header("Set-Cookie", "token=; Path=/; Max-Age=0;")
	c.JSON(http.StatusOK, gin.H{
		"isValid": true,
		"result":  true,
	})
}