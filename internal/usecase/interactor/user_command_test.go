package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory/mock_factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
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
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterは呼ばれない
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(0)
	mockUserRepository.EXPECT().Save(gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.AddUser(command.UserAddInputData{}, mockContext)
}

// 異常系のテスト: UserFactoryでエラーが発生した時
func TestUserUseCaseInteractor_AddUser_NegativeFactoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Return(entity.User{}, errors.New("error")).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Times(0)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.AddUser(command.UserAddInputData{}, mockContext)
}

// 異常系のテスト: UserRepositoryでエラーが発生した時
func TestUserUseCaseInteractor_AddUser_NegativeRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// error presenterが呼ばれる
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Return(entity.User{}, nil).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Return(errors.New("error")).Times(1)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockErrorPresenter, mockUserRepository, mockUserFactory)
	useCase.AddUser(command.UserAddInputData{}, mockContext)
}