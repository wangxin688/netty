package auth

import (
	"netty/core"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGenerateTokenResponse(t *testing.T) {
	userId := 1
	username := "testuser"

	// Test generating token response for a valid user ID and username
	t.Run("Generate token response", func(t *testing.T) {
		tokenResponse, err := GenerateTokenResponse(userId, username)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if tokenResponse == nil {
			t.Error("Expected token response, but got nil")
		}
	})

	// Test error handling when creating access token fails
	t.Run("Error creating access token", func(t *testing.T) {
		core.Settings.AccessTokenExpireMinutes = -1 // Set invalid expire minutes to force an error
		_, err := GenerateTokenResponse(userId, username)
		if err == nil {
			t.Error("Expected error when creating access token with invalid expire minutes")
		}
		core.Settings.AccessTokenExpireMinutes = 60 // Reset expire minutes
	})

	// Test error handling when creating refresh token fails
	t.Run("Error creating refresh token", func(t *testing.T) {
		core.Settings.RefreshTokenExpireMinutes = -1 // Set invalid expire minutes to force an error
		_, err := GenerateTokenResponse(userId, username)
		if err == nil {
			t.Error("Expected error when creating refresh token with invalid expire minutes")
		}
		core.Settings.RefreshTokenExpireMinutes = 1440 // Reset expire minutes
	})
}

func TestPasswordHashed(t *testing.T) {
	password := "testpassword"

	t.Run("Hashing a password successfully", func(t *testing.T) {
		hashedPassword := PasswordHashed(password)

		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
			t.Errorf("Expected hashed password to match original password, but they did not match")
		}
	})

	t.Run("Hashing an empty password", func(t *testing.T) {
		emptyPassword := ""
		hashedEmptyPassword := PasswordHashed(emptyPassword)

		if err := bcrypt.CompareHashAndPassword([]byte(hashedEmptyPassword), []byte(emptyPassword)); err == nil {
			t.Errorf("Expected hashing an empty password to return an error, but it did not")
		}
	})
}

func TestPasswordVerify(t *testing.T) {
	hashedPassword := "$2a$10$Rl4iJlJG.8sXpB6Bpr0W4eDZPm2eXf0UrVof3b9b5VuL91nRc9nAu"
	password := "secret"

	// Test case for matching password and hashed password
	if !PasswordVerify(hashedPassword, password) {
		t.Errorf("Expected matching password and hashed password to return true")
	}

	// Test case for non-matching password and hashed password
	if PasswordVerify(hashedPassword, "wrongpassword") {
		t.Errorf("Expected non-matching password and hashed password to return false")
	}
}

func TestRandomTokenString(t *testing.T) {
	t.Run("Generate token of length 0", func(t *testing.T) {
		token := RandomTokenString(0)
		if len(token) != 0 {
			t.Errorf("Expected token length to be 0, got %d", len(token))
		}
	})

	t.Run("Generate token of length 10", func(t *testing.T) {
		token := RandomTokenString(10)
		if len(token) != 20 { // Since each byte is encoded as 2 characters in hex
			t.Errorf("Expected token length to be 20, got %d", len(token))
		}
	})

	t.Run("Generate token of length 20", func(t *testing.T) {
		token := RandomTokenString(20)
		if len(token) != 40 { // Since each byte is encoded as 2 characters in hex
			t.Errorf("Expected token length to be 40, got %d", len(token))
		}
	})
}
