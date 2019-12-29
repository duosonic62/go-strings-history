package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary/mock_outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestErrorUseCaseInteractor_BadRequestError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewBadRequestErrorUseCase(mockErrorPresenter)
	useCase.BadRequestError(mockContext, errors.New("error"))
}