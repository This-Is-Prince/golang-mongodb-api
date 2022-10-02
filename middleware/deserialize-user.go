package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/This-Is-Prince/golang-mongodb-api/config"
	"github.com/This-Is-Prince/golang-mongodb-api/services"
	"github.com/This-Is-Prince/golang-mongodb-api/utils"
	"github.com/gin-gonic/gin"
)

func DeserializeUser(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := userService.FindUserById(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no longer exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
