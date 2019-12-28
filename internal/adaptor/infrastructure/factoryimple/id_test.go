package factoryimple

import (
	"testing"
)

func TestIDFactoryImpl_Gen(t *testing.T) {
	factory := NewIdFactory()
	actual, err := factory.Gen()

	if err != nil {
		t.Error(err)
	}

	if len(actual) != 36 {
		t.Error("Expected: 36 characters")
	}
}