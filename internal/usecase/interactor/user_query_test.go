package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary/mock_outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserQueryUseCaseInteracter_Show_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepository := mock_repository.NewMockUserQueryRepository(ctrl)
	mockPresenter := mock_outputboundary.NewMockUserQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)

	mockRepository.EXPECT().Find(gomock.Any()).Return(output.UserOutput{Name: "mock_name"}, nil).Times(1)
	mockPresenter.EXPECT().OutputUser(gomock.Any(), gomock.Any()).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(0)
	mockContext := mock_input.NewMockContext(ctrl)
	mockToken, _ := valueobject.NewAuthorizationToken("mock_token")

	useCase := NewUserQueyUseCase(mockRepository, mockPresenter, mockErrorPresenter)
	useCase.Show(mockToken, mockContext)
}

func TestUserQueryUseCaseInteracter_Show_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepository := mock_repository.NewMockUserQueryRepository(ctrl)
	mockPresenter := mock_outputboundary.NewMockUserQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)

	mockRepository.EXPECT().Find(gomock.Any()).Return(output.UserOutput{Name: "mock_name"}, errors.New("error")).Times(1)
	mockPresenter.EXPECT().OutputUser(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockContext := mock_input.NewMockContext(ctrl)
	mockToken, _ := valueobject.NewAuthorizationToken("mock_token")

	useCase := NewUserQueyUseCase(mockRepository, mockPresenter, mockErrorPresenter)
	useCase.Show(mockToken, mockContext)
}