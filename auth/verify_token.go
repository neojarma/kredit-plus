package auth

import (
	"fmt"
	"kredit_plus/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var JwtAccessSecret = []byte(os.Getenv("JWT_ACCESS_SECRET"))
var JwtRefreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authorization header"})
		}

		tokenString := ""
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authorization header format"})
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtAccessSecret, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		claims, ok := token.Claims.(*models.JWTClaims)
		if !ok || claims == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid claims"})
		}

		c.Set("user_id", claims.UserID)

		return next(c)
	}
}

func GenerateAccessToken(userID string) (string, error) {
	claims := models.JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtAccessSecret)
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := models.JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtRefreshSecret)
}

func CheckRefreshToken(refreshToken string) (*models.LoginResponse, error) {
	claims, err := isValid(refreshToken)
	if err != nil {
		return nil, echo.ErrUnauthorized
	}

	newAccessToken, err := GenerateAccessToken(claims.UserID)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ID:            claims.Id,
		AccessToken:   newAccessToken,
		RefreshTokens: refreshToken,
	}, nil
}

func isValid(tokenString string) (*models.JWTClaims, error) {
	claims := &models.JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtRefreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
