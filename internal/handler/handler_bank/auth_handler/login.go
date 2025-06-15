package authhandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/riyanathariq/taskify-api/internal/common"
	"github.com/riyanathariq/taskify-api/internal/dependencies"
	"github.com/riyanathariq/taskify-api/internal/models"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func LoginHandler(c *gin.Context) {
	var (
		req = LoginRequest{}
		ctx = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Get user
	user, err := dependencies.New().Repository.User.DetailByUsername(ctx, req.Username)
	if err != nil || !common.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// JWT Claims
	cfg := dependencies.New().Config
	now := time.Now()

	accessTokenExp := now.Add(180 * time.Minute)

	tokenID := uuid.New().String()

	accessTokenClaims := jwt.MapClaims{
		"sub":      tokenID,
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      accessTokenExp.Unix(),
		"iss":      cfg.AppName,
		"iat":      now.Unix(),
	}

	// Sign tokens
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign access token"})
		return
	}

	// Save to DB
	tokenModel := &models.Token{
		ID:              tokenID,
		UserID:          user.ID,
		AccessToken:     accessToken,
		AccessExpiresAt: accessTokenExp,
		Revoked:         false,
		CreatedAt:       now,
	}
	if err := dependencies.New().Repository.Oauth.SaveToken(ctx, tokenModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store token"})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken: accessToken,
	})
}
