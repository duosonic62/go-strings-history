package controller

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary/mock_inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserControllerImpl_CreateUser_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockBadRequestUseCase := mock_inputboundary.NewMockBadRequestUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(nil).Times(1)
	mockBadRequestUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockUserCommandUseCase.EXPECT().AddUser(gomock.Any(), gomock.Any()).Times(1)

	controller := NewUserController(mockUserCommandUseCase, mockBadRequestUseCase)
	controller.CreateUser(mockContext)
}

func TestUserControllerImpl_CreateUser_NegativeBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockBadRequestUseCase := mock_inputboundary.NewMockBadRequestUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(errors.New("error")).Times(1)
	mockBadRequestUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(1)
	mockUserCommandUseCase.EXPECT().AddUser(gomock.Any(), gomock.Any()).Times(0)

	controller := NewUserController(mockUserCommandUseCase, mockBadRequestUseCase)
	controller.CreateUser(mockContext)
}