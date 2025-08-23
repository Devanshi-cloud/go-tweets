package user

import (
	"context"
	"errors"
	"go-tweets/internal/dto"
	"go-tweets/internal/model"
	"go-tweets/pkg/jwt"
	"go-tweets/pkg/refreshtoken"
	"net/http"
	"time"
)

func (s *userService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error) {  
	// check user existence
	userExist, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err // Internal server error
	}
	if userExist == nil {
		return "", "", http.StatusNotFound, errors.New("user not found") // User not found
	}
	// Get the refresh token by user ID
	refreshTokenExists, err := s.userRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		return "", "", http.StatusInternalServerError, err // Error fetching refresh token
	}
	
	if refreshTokenExists == nil {
		return "", "", http.StatusUnauthorized, errors.New("refresh token was expired") // Refresh token not found
	}

	//check refresh token validity
	if req.RefreshToken != refreshTokenExists.RefreshToken {
		return "", "", http.StatusUnauthorized, errors.New("refresh token not found") // Invalid refresh token
	}

	//generate a new refresh token if the matched
	token, err := jwt.CreateToken(userID, userExist.Username, s.cfg.SecretJwt)
	if err != nil {
		return "", "", http.StatusInternalServerError, err // Error creating token
	}

	err = s.userRepo.DeleteRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return "", "", http.StatusInternalServerError, err // Error deleting old refresh token
	}
	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err // Error creating new refresh token
	}

	now := time.Now()
	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:      userID,
		RefreshToken: refreshToken, // New refresh token
		CreatedAt:  now,
		UpdatedAt: now,
		ExpiredAt: now.Add(24 * time.Hour), // Set expiration time
	})
	if err != nil {
		return "", "", http.StatusInternalServerError, err // Error storing new refresh token
	}

	return token, refreshToken, http.StatusOK, nil // Return new access token and refresh token
}