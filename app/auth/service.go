package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"netty/core"

	"golang.org/x/crypto/bcrypt"
)

// GenerateTokenResponse generates an access and refresh token for the given user ID and username
func GenerateTokenResponse(userId int, username string) (*core.AccessToken, error) {
	// Create an access token
	accessToken, expiresAt, issuedAt, err := core.CreateAccessToken(
		userId, username, false, core.Settings.AccessTokenExpireMinutes)
	if err != nil {
		return nil, err
	}

	// Create a refresh token
	refreshToken, refreshExpiresAt, refreshIssuedAt, err := core.CreateAccessToken(
		userId, username, true, core.Settings.RefreshTokenExpireMinutes)
	if err != nil {
		return nil, err
	}
	return &core.AccessToken{
		TokenType: core.AuthorizationBearer,
		AccessToken: accessToken,
		ExpiresAt: expiresAt,
		IssuedAt: issuedAt,
		RefreshToken: refreshToken,
		RefreshTokenExpiresAt: refreshExpiresAt,
		RefreshTokenIssuedAt: refreshIssuedAt,
	}, nil
}



// PasswordHashed generates a hashed password from the given plain text password.
//
// password: a string representing the plain text password.
// string: returns the hashed password as a string.
func PasswordHashed(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashedPassword)
}

// PasswordVerify checks if the provided password matches the hashed password.
//
// hashedPassword string, password string. bool
func PasswordVerify(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// RandomTokenString generates a random token string of specified length.
//
// length int - the length of the token string
// string - the randomly generated token string
func RandomTokenString(length int) (token string) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Errorf("unable to generate random token: %w", err))
	}
	token = hex.EncodeToString(b)
	return
}
