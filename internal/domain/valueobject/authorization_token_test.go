package valueobject

import "testing"

func TestNewAuthorizationToken_Positive(t *testing.T) {
	_, err := NewAuthorizationToken("mock_token")
	if err != nil {
		t.Error("Expected: Error is nil")
	}
}

func TestNewAuthorizationToken_Negative(t *testing.T) {
	_, err := NewAuthorizationToken("") // token validation error
	if err == nil {
		t.Error("Expected: error is not nil")
	}
}
