package factory

import (
	"github.com/friendsofgo/errors"
	"testing"
)

// 正常系のIDFactoryのモック
type mockIDFactory struct {}
func (m mockIDFactory) Gen() (string, error) {
	return "mock_id", nil
}

// 正常系のTokenFactoryのモック
type mockTokenFactory struct {}
func (m mockTokenFactory) Gen() (string, error) {
	return "mock_token", nil
}

// 異常系のIDFactoryのモック
type mockErrorIDFactory struct {}
func (m mockErrorIDFactory) Gen() (string, error) {
	return "", errors.New("id generate error")
}

// 異常系のTokenFactoryのモック
type mockErrorTokenFactory struct {}
func (m mockErrorTokenFactory) Gen() (string, error) {
	return "", errors.New("token generate error")
}

func TestUserFactory_NewUser(t *testing.T) {
	factory := NewUserFactory(mockIDFactory{}, mockTokenFactory{})
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
	factory := NewUserFactory(mockErrorIDFactory{}, mockTokenFactory{})
	_, err := factory.NewUser("mock_name", "mock_uid")

	if err == nil {
		t.Error("Expected: error must not be nil")
	}
}

func TestUserFactory_NewUser_Negative_ErrorToken(t *testing.T) {
	factory := NewUserFactory(mockIDFactory{}, mockErrorTokenFactory{})
	_, err := factory.NewUser("mock_name", "mock_uid")

	if err == nil {
		t.Error("Expected: error must not be nil")
	}
}