package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary/mock_inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserControllerImpl_Show_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserQueryService := mock_inputboundary.NewMockUserQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockUserQueryService.EXPECT().Show(gomock.Any(), gomock.Any()).Times(1)

	controller := NewUserQueryController(mockUserQueryService, mockErrorUseCase)
	controller.Show(mockContext)
}

func TestUserQueryControllerImpl_Show_NegativeNoAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserQueryService := mock_inputboundary.NewMockUserQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1) // error token
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockUserQueryService.EXPECT().Show(gomock.Any(), gomock.Any()).Times(0)

	controller := NewUserQueryController(mockUserQueryService, mockErrorUseCase)
	controller.Show(mockContext)
}