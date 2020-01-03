package factory

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/factory/mock_factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserFactory_NewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)

	mockIDFactory.EXPECT().Gen().Return("mock_id", nil)
	mockTokenFactory.EXPECT().Gen().Return("mock_token", nil)

	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	factory := NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	actual, err := factory.NewUser("mock_name", "mock_uid")

	if err != nil {
		t.Error(err)
	}

	if actual.ID != "mock_id" {
		t.Errorf("Expected: mock_id, Actual %s\n", actual.ID)
	}

	if actual.Name != "mock_name" {
		t.Errorf("Expected: mock_name, Actual %s\n", actual.Name)
	}

	if actual.UID != "mock_uid" {
		t.Errorf("Expected: mock_uid, Actual %s\n", actual.UID)
	}

	if actual.Token != "mock_token" {
		t.Errorf("Expected: mock_token, Actual %s\n", actual.Token)
	}
}

func TestUserFactory_NewUser_Negative_ErrorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)

	mockIDFactory.EXPECT().Gen().Return("mock_id", errors.New("error"))
	// エラーで早期リターンしているので呼ばれない
	// mockTokenFactory.EXPECT().Gen().Return("mock_token", nil)

	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	factory := NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	_, err := factory.NewUser("mock_name", "mock_uid")

	if err == nil {
		t.Error("Expected: error must not be nil")
	}
}

func TestUserFactory_NewUser_Negative_ErrorToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)

	mockIDFactory.EXPECT().Gen().Return("mock_id", nil)
	mockTokenFactory.EXPECT().Gen().Return("mock_token", errors.New("error"))

	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	factory := NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	_, err := factory.NewUser("mock_name", "mock_uid")

	if err == nil {
		t.Error("Expected: error must not be nil")
	}
}