package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// Middleware để kiểm tra xác thực người dùng qua session
var store = sessions.NewCookieStore([]byte("secret-key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy session
		session, _ := store.Get(c.Request, "session-name")
		userID, ok := session.Values["user_id"]
		if !ok || userID == nil {
			// Nếu không có user_id trong session, trả về lỗi 401 (Unauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Nếu đã xác thực, cho phép tiếp tục với yêu cầu
		c.Next()
	}
}
