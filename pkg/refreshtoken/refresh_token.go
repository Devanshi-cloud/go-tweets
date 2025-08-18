package refreshtoken

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() (string, error) {
	// Generate a random 18-byte token
	tokenBytes := make([]byte, 18)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err // Return error if token generation fails
	}

	// Convert the byte slice to a hexadecimal string
	return hex.EncodeToString(tokenBytes), nil // Return the generated refresh token
}
