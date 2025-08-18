package user

import (
    "context"
    "errors"
    "go-tweets/internal/dto"
    "go-tweets/internal/model"
    "go-tweets/pkg/jwt"
    "net/http"
    "time"

    "go-tweets/pkg/refreshtoken"

    "golang.org/x/crypto/bcrypt"
)

func(s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
    // check if the user exists
    userExist, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, "")
    if err != nil {
        return "", "", http.StatusInternalServerError, err
    }

    if userExist == nil {
        return "", "", http.StatusNotFound, errors.New("wrong email or password") // User not found
    }

    err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
    if err != nil {
        return "", "", http.StatusNotFound, errors.New("wrong email or password") // Password mismatch
    }

    //generate access token and refresh token
    token, err := jwt.CreateToken(userExist.ID, userExist.Username, s.cfg.SecretJwt)
    if err != nil {
        return "", "", http.StatusInternalServerError, err // Error creating token
    }

    //get refresh token if exists
    now := time.Now()
    refreshTokenExist, err := s.userRepo.GetRefreshToken(ctx, userExist.ID, now)
    if err != nil {
        return "", "", http.StatusInternalServerError, err // Error fetching refresh token
    }

    if refreshTokenExist != nil {
        return token, refreshTokenExist.RefreshToken, http.StatusOK, nil // Return existing refresh token
    } // Close the if block

    //generate & share refresh token
    refreshToken, err := refreshtoken.GenerateRefreshToken()
    if err != nil {
        return "", "", http.StatusInternalServerError, err // Error creating refresh token
    }

    err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
        UserID:      userExist.ID,
        RefreshToken: refreshToken,
        CreatedAt:   now,
        UpdatedAt: now,
        ExpiredAt:   time.Now().Add(7 * 24 * time.Hour), // Set expiration to 7 days
    })
    if err != nil {
        return "", "", http.StatusInternalServerError, err // Error storing refresh token
    }

    // Save the refresh token in the database
    return token, refreshToken, http.StatusOK, nil // Return new refresh token

}