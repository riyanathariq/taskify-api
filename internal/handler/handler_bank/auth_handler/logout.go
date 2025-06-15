package authhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/riyanathariq/taskify-api/internal/consts"
	"github.com/riyanathariq/taskify-api/internal/dependencies"
	"net/http"
)

func LogoutHandler(c *gin.Context) {
	ctx := c.Request.Context()

	tokenID := ctx.Value(consts.TokenID).(string)

	_ = dependencies.New().Repository.Oauth.RevokeToken(ctx, tokenID)

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
