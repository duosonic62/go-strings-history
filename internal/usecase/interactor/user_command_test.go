package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory/mock_factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary/mock_outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

// 正常系のテスト
func TestUserUseCaseInteractor_AddUser_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterは呼ばれない
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(0)
	mockUserRepository.EXPECT().Save(gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Times(1).Return(user(), nil)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Add(command.UserAddInputData{}, mockContext)
}

// 異常系のテスト: UserFactoryでエラーが発生した時
func TestUserUseCaseInteractor_AddUser_NegativeFactoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Times(0)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Add(command.UserAddInputData{}, mockContext)
}

// 異常系のテスト: UserRepositoryでエラーが発生した時
func TestUserUseCaseInteractor_AddUser_NegativeRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Return(user(), nil).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Return(errors.New("error")).Times(1)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Add(command.UserAddInputData{}, mockContext)
}

func TestUserUseCaseInteractor_Edit_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterは呼ばれない
	mockUserPresenter.EXPECT().OutputEditUser(gomock.Any(), gomock.Any()).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(0)
	mockUserFactory.EXPECT().Find(gomock.Any()).Times(1).Return(user(), nil)
	mockUserRepository.EXPECT().Edit(gomock.Any()).Times(1)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Edit(token(), command.UserEditInputData{}, mockContext)
}

func TestUserUseCaseInteractor_Edit_Negative_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockUserPresenter.EXPECT().OutputEditUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserFactory.EXPECT().Find(gomock.Any()).Times(1).Return(nil, errors.New("error"))
	mockUserRepository.EXPECT().Edit(gomock.Any()).Times(0)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Edit(token(), command.UserEditInputData{}, mockContext)
}

func TestUserUseCaseInteractor_Edit_Negative_FailEdit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockUserPresenter.EXPECT().OutputEditUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserFactory.EXPECT().Find(gomock.Any()).Times(1).Return(user(), nil)
	mockUserRepository.EXPECT().Edit(gomock.Any()).Times(1).Return(errors.New("error"))

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Edit(token(), command.UserEditInputData{}, mockContext)
}

func TestUserUseCaseInteractor_Delete_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterは呼ばれない
	mockUserPresenter.EXPECT().OutputDeleteUser(gomock.Any()).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(0)
	mockUserRepository.EXPECT().Delete(gomock.Any()).Times(1)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Delete(token(), mockContext)
}

func TestUserUseCaseInteractor_Delete_Negative_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserCommandRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserRepository.EXPECT().Delete(gomock.Any()).Times(1).Return(errors.New("error"))

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.Delete(token(), mockContext)
}

// テストヘルパーメソッド
func user() *entity.User {
	token, _ := valueobject.NewAuthorizationToken("mock_token")
	user, _ := entity.NewUser("mock_id", "mock_name",  "mock_uid", token)
	return user
}

func token() valueobject.AuthorizationToken {
	token, _ := valueobject.NewAuthorizationToken("mock_token")
	return token
}