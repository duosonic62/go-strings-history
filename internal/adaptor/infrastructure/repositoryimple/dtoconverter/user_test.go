package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"testing"
)

func TestConvertUser(t *testing.T) {
	user := &models.Member{Name: "mock_name"}
	actual := ConvertUser(user)
	expected := output.UserOutput{Name: "mock_name"}

	if actual.Name != expected.Name {
		t.Errorf("Expected: mock_name, Actual %s\n", actual.Name)
	}
}
