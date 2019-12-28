package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"testing"
)

// presenterのモック
type mockUserPresenter struct{}

func (mockUserPresenter) OutputAddUser(data output.UserAddOutputData, ctx input.Context) {
	// 何もしない
}

// repositoryのモック
type mockUserRepository struct{}

func (mockUserRepository) Save(entity.User) {
	// 何もしない
}

// 正常系のUserFactoryのモック
type mockUserFactory struct{}

func (mockUserFactory) NewUser(name string, uid string) (entity.User, error) {
	return entity.User{
		ID:    "mock_id",
		Name:  "mock_name",
		UID:   "mock_uid",
		Token: "mock_token",
	}, nil
}

// 異常系のUserFactoryのモック
type mockUserErrorFactory struct{}

func (mockUserErrorFactory) NewUser(name string, uid string) (entity.User, error) {
	return entity.User{}, errors.New("mock error")
}

// contextのモック
type mockContext struct{}

func (mockContext) Param(string) string    { return "mock_param" }
func (mockContext) Bind(interface{}) error { return nil }
func (mockContext) Status(int)             {}
func (mockContext) JSON(int, interface{})  {}

// TODO: 正常系のテスト
func TestUserUseCaseInteractor_AddUser_Positive(t *testing.T) {
	useCase := NewUserCommandUseCase(mockUserPresenter{}, mockUserRepository{}, mockUserFactory{})
	useCase.AddUser(command.UserAddInputData{}, mockContext{})
}

// TODO: 異常系のテスト
func TestUserUseCaseInteractor_AddUser_Negative(t *testing.T) {
	useCase := NewUserCommandUseCase(mockUserPresenter{}, mockUserRepository{}, mockUserErrorFactory{})
	useCase.AddUser(command.UserAddInputData{}, mockContext{})
}
