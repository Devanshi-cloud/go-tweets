package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	// Implement JWT token creation logic here
	// This function should create a JWT token with the user ID and username as claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"id": id,
			"username": username,
			"exp": time.Now().Add(60 * time.Minute).Unix(), // Token expiration time
		},
	)
	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)

	return tokenStr, err
}

func ValidateToken(tokenStr, secretKey string, withClaimValidation bool) (int64, string, error) {
	var (
		key = []byte(secretKey)
		claims = jwt.MapClaims{}
		token *jwt.Token
		err error
	)

	if withClaimValidation {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}

	if err != nil {
		return 0, "", err // Return error if token parsing fails
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token") // Return error if token is invalid
	}

	userID := int64(claims["id"].(float64))
	username := claims["username"].(string)
	
	return userID, username, nil // Return user ID and username if token is valid
}

