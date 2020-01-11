package controller

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary/mock_inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserCommandControllerImpl_Create_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(nil).Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockUserCommandUseCase.EXPECT().Add(gomock.Any(), gomock.Any()).Times(1)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Create(mockContext)
}

func TestUserCommandControllerImpl_Create_NegativeBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(errors.New("error")).Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(1)
	mockUserCommandUseCase.EXPECT().Add(gomock.Any(), gomock.Any()).Times(0)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Create(mockContext)
}

func TestUserCommandControllerImpl_Edit_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(nil).Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockUserCommandUseCase.EXPECT().Edit(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Edit(mockContext)
}

func TestUserCommandControllerImpl_Edit_NegativeBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(errors.New("error")).Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(0)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(1)
	mockUserCommandUseCase.EXPECT().Edit(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Edit(mockContext)
}

func TestUserCommandControllerImpl_Edit_NegativeUnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(nil).Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1) // Headerに不正な値をセット
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockUserCommandUseCase.EXPECT().Edit(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Edit(mockContext)
}

func TestUserCommandControllerImpl_Delete_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockUserCommandUseCase.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(1)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Delete(mockContext)
}

func TestUserCommandControllerImpl_Delete_NegativeUnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserCommandUseCase := mock_inputboundary.NewMockUserCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1) // Headerに不正な値をセット
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockUserCommandUseCase.EXPECT().Delete(gomock.Any(), gomock.Any()).Times(0)

	controller := NewUserController(mockUserCommandUseCase, mockErrorUseCase)
	controller.Delete(mockContext)
}
