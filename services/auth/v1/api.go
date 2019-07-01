package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login : 登陆
func (a *Auth) Login(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
