package routes

import (
	"net/http"

	"L-cart/controllers/org"
	"L-cart/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// 認証不要のエンドポイント
	r.POST("/org/login", org.HandleOrgLogin)

	// 認証が必要なエンドポイント
	authorized := r.Group("/org")
	authorized.Use(middleware.OrgAuthMiddleware())
	{
		authorized.POST("/logout", org.HandleLogout)
	}
}
