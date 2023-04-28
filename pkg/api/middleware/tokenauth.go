package middleware

import (
	"fmt"
	"net/http"
	"time"

	"shiftsync/pkg/auth"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser(ctx *gin.Context) {
	authtoken(ctx, "employee-cookie")
}

func authtoken(ctxt *gin.Context, user string) {
	token, err := ctxt.Cookie(user)
	fmt.Println(token)

	if err != nil || token == "" {
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status code ": 401,
			"msg":          "Unauthorized User Please Login",
		})
		return
	}

	claims, err := auth.ValidateTokens(token)
	fmt.Println(claims)

	if err != nil {
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
			"err":        err,
		})
		return
	}

	if time.Now().Unix() > claims.ExpiresAt {
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "User Need Re-Login time expired",
		})
		return
	}

	// claim the userId and set it on context
	ctxt.Set("employeeId", claims.Id)
}
