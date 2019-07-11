package middleware

import (
	"github.com/asynccnu/grade_service_v2/handler"
	"github.com/asynccnu/grade_service_v2/pkg/errno"
	"github.com/asynccnu/grade_service_v2/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if err := token.ParseRequest(c); err != nil {
			handler.SendUnAuth(c, errno.ErrAuthorizationInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
