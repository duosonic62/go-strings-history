package service

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory/mock_factory"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAuthorizationServiceImpl_Authorized_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockUserFactory.EXPECT().Find(gomock.Any()).Return(user(), nil).Times(1)

	service := NewAuthorizationService(mockUserFactory)
	_, err := service.Authorized(token())

	if err != nil {
		t.Error("Expected: error is nil")
	}
}

func TestAuthorizationServiceImpl_Authorized_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockUserFactory.EXPECT().Find(gomock.Any()).Return(nil, errors.New("error")).Times(1)

	service := NewAuthorizationService(mockUserFactory)
	actual, err := service.Authorized(token())

	if actual != nil && err == nil {
		t.Error("Expected: error is not nil")
	}
}


// テストヘルパーメソッド
func user() *entity.User {
	token, _ := valueobject.NewAuthorizationToken("mock_token")
	user, _ := entity.NewUser("mock_id", "mock_name",  "mock_uid", token)
	return user
}

func token() *valueobject.AuthorizationToken {
	token, _ := valueobject.NewAuthorizationToken("mock_token")
	return &token
}