package jwt_helper

import (
	"sigap-sultan-be/src/app/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("TokenPangan")

// Struct to represent JWT claims
type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(claimData *models.LoginResponse) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claimData.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Issuer:    "AppPangan",
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claimData)
	token, err := tokenClaims.SignedString(mySigningKey)

	return token, err
}

func ParseToken(token string) (*models.LoginResponse, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &models.LoginResponse{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*models.LoginResponse); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
