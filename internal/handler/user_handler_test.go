package handler

import (
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yockii/Tianshu/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

// TestPasswordHashAndCompare tests bcrypt password hashing and comparison
func TestPasswordHashAndCompare(t *testing.T) {
	password := "Test@123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to generate hash: %v", err)
	}
	// Correct password should match
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		t.Errorf("Expected password to match: %v", err)
	}
	// Wrong password should not match
	if err := bcrypt.CompareHashAndPassword(hash, []byte("wrongpass")); err == nil {
		t.Errorf("Expected password mismatch error, got nil")
	}
}

// TestJWTGenerateAndParse tests JWT creation and parsing using config
func TestJWTGenerateAndParse(t *testing.T) {
	// Initialize config to load JWT secret
	if err := config.InitConfig(""); err != nil {
		t.Fatalf("Failed to init config: %v", err)
	}
	secret := config.Cfg.JWT.Secret
	claimsOrig := jwt.MapClaims{
		"userId":   1,
		"tenantId": 100,
		"exp":      time.Now().Add(time.Hour * time.Duration(config.Cfg.JWT.ExpireHours)).Unix(),
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsOrig).SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.Fatalf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}
	if !token.Valid {
		t.Errorf("Token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("Failed to get claims")
	}
	if claims["userId"] != claimsOrig["userId"] {
		t.Errorf("Expected userId %v, got %v", claimsOrig["userId"], claims["userId"])
	}
	if claims["tenantId"] != claimsOrig["tenantId"] {
		t.Errorf("Expected tenantId %v, got %v", claimsOrig["tenantId"], claims["tenantId"])
	}
}
