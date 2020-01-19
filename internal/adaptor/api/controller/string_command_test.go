package controller

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary/mock_inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestStringCommandControllerImpl_Create_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringCommandUseCase := mock_inputboundary.NewMockStringCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)

	mockContext.EXPECT().Bind(gomock.Any()).Return(nil).Times(1)
	mockStringCommandUseCase.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

	controller := NewStringCommandController(mockStringCommandUseCase, mockErrorUseCase)
	controller.Create(mockContext)
}

func TestStringCommandControllerImpl_Create_Negative_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringCommandUseCase := mock_inputboundary.NewMockStringCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(errors.New("error")).Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(0)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(1)
	mockStringCommandUseCase.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewStringCommandController(mockStringCommandUseCase, mockErrorUseCase)
	controller.Create(mockContext)
}

func TestStringCommandControllerImpl_Create_Negative_UnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringCommandUseCase := mock_inputboundary.NewMockStringCommandUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Bind(gomock.Any()).Return(nil).Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockStringCommandUseCase.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewStringCommandController(mockStringCommandUseCase, mockErrorUseCase)
	controller.Create(mockContext)
}