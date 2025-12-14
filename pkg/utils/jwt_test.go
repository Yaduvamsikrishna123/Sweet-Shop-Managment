package utils

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	token, err := GenerateToken(1, "user")
	if err != nil {
		t.Errorf("GenerateToken failed: %v", err)
	}
	if token == "" {
		t.Error("GenerateToken returned empty string")
	}
}
