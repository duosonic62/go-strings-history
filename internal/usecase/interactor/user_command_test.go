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
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// それぞれ一回づつ呼ばれる
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewUserCommandUseCase(mockUserPresenter, mockUserRepository, mockUserFactory)
	useCase.AddUser(command.UserAddInputData{}, mockContext)
}

// TODO: 異常系のテスト
func TestUserUseCaseInteractor_AddUser_Negative(t *testing.T) {
	// エラーハンドリングがまだなので一旦スキップ
	t.Skip("TODO: Error Handling...")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserPresenter := mock_outputboundary.NewMockUserCommandPresenter(ctrl)
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockUserFactory := mock_factory.NewMockUserFactory(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	// それぞれ一回づつ呼ばれる
	mockUserPresenter.EXPECT().OutputAddUser(gomock.Any(), gomock.Any()).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Times(1)
	mockUserFactory.EXPECT().NewUser(gomock.Any(), gomock.Any()).Return(entity.User{}, errors.New("error")).Times(1)


	useCase := NewUserCommandUseCase(mockUserPresenter, mockUserRepository, mockUserFactory)
	useCase.AddUser(command.UserAddInputData{}, mockContext)
}
