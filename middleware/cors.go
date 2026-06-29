package middleware

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// HKF-fix(2026-06-29): token 鉴权 API 不需要 credentials。原 AllowAllOrigins+AllowCredentials=true
	// 会反射来源+带凭证→可被 CSRF(console 有会话 cookie)。关掉 credentials,gin-cors 发 Allow-Origin:* 无凭证。
	config.AllowCredentials = false
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	return cors.New(config)
}

func PoweredBy() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-New-Api-Version", common.Version)
		c.Next()
	}
}
