package factoryimple

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory/mock_factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserFactory_NewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)
	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	mockIDFactory.EXPECT().Gen().Return("mock_id", nil)
	mockTokenFactory.EXPECT().Gen().Return("mock_token", nil)

	factory := NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	actual, err := factory.NewUser("mock_name", "mock_uid")

	if err != nil {
		t.Error(err)
	}

	if actual.GetID() != "mock_id" {
		t.Errorf("Expected: mock_id, Actual %s\n", actual.GetID())
	}

	if actual.GetName() != "mock_name" {
		t.Errorf("Expected: mock_name, Actual %s\n", actual.GetName())
	}

	if actual.GetUID() != "mock_uid" {
		t.Errorf("Expected: mock_uid, Actual %s\n", actual.GetUID())
	}

	if actual.GetToken().GetToken() != "mock_token" {
		t.Errorf("Expected: mock_token, Actual %s\n", actual.GetToken().GetToken())
	}
}

func TestUserFactory_NewUser_Negative_ErrorID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)
	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	mockIDFactory.EXPECT().Gen().Return("mock_id", errors.New("error"))
	// エラーで早期リターンしているので呼ばれない
	// mockTokenFactory.EXPECT().Gen().Return("mock_token", nil)

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
	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	mockIDFactory.EXPECT().Gen().Return("mock_id", nil)
	mockTokenFactory.EXPECT().Gen().Return("mock_token", errors.New("error"))


	factory := NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	_, err := factory.NewUser("mock_name", "mock_uid")

	if err == nil {
		t.Error("Expected: error must not be nil")
	}
}

func TestUserFactoryImpl_Find_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)
	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	mockRepository.EXPECT().Find(gomock.Any()).Times(1).Return(user(), nil)

	factory :=  NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	_, err := factory.Find(token())
	if err != nil {
		t.Error("Expected: error must be nil")
	}
}

func TestUserFactoryImpl_Find_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIDFactory := mock_factory.NewMockIDFactory(ctrl)
	mockTokenFactory := mock_factory.NewMockTokenFactory(ctrl)
	mockRepository := mock_repository.NewMockUserCommandRepository(ctrl)

	mockRepository.EXPECT().Find(gomock.Any()).Times(1).Return(nil, errors.New("error"))

	factory :=  NewUserFactory(mockIDFactory, mockTokenFactory, mockRepository)
	_, err := factory.Find(token())
	if err == nil {
		t.Error("Expected: error must not be nil")
	}
}

func user() *entity.User {
	token, _ := valueobject.NewAuthorizationToken("mock_token")
	user, _ := entity.NewUser("mock_id", "mock_name",  "mock_uid", token)
	return user
}

func token() *valueobject.AuthorizationToken {
	token, _ := valueobject.NewAuthorizationToken("mock_token")
	return &token
}