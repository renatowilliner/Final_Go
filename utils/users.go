package utils

import (
	"examen/clients/responses"

	"github.com/gin-gonic/gin"
)

const (
	RolAdministrador = "ADMIN"
	RolUsuario       = "USUARIO"
	
)

func SetUserInContext(c *gin.Context, user *responses.UserInfo) {
	c.Set("UserInfo", user)
}

func GetUserInfoFromContext(c *gin.Context) *responses.UserInfo {
	userInfo, _ := c.Get("UserInfo")

	user, _ := userInfo.(*responses.UserInfo)

	return user
}
