package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nicitapa/firstProgect/internal/configs"
	"github.com/nicitapa/firstProgect/internal/models"
	"github.com/nicitapa/firstProgect/pkg"
	"strings"
)

func (ctrl *Controller) extractTokenFromHeader(c *gin.Context, headerKey string) (string, error) {
	header := c.GetHeader(headerKey)

	if header == "" {
		return "", errors.New("empty authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty token")
	}

	return headerParts[1], nil
}

func (ctrl *Controller) generateNewTokenPair(userID int, userRole models.Role) (string, string, error) {
	// сгенерировать токен (браслет)
	accessToken, err := pkg.GenerateToken(userID,
		configs.AppSettings.AuthParams.AccessTokenTllMinutes,
		userRole, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateToken(userID,
		configs.AppSettings.AuthParams.RefreshTokenTllDays,
		userRole, true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
