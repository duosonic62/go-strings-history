package factoryimple

import (
	"testing"
)

func TestTokenFactoryImpl_Gen(t *testing.T) {
	factory := NewTokenFactory()
	actual, err := factory.Gen()

	if err != nil {
		t.Error(err)
	}

	if len(actual) != 36 {
		t.Error("Expected: 36 characters")
	}
}