package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary/mock_inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestStringQueryController_GetGuitarString_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringQueryUseCase := mock_inputboundary.NewMockStringQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Param(gomock.Any()).Return("mock_param").Times(1)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(0)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockStringQueryUseCase.EXPECT().GetGuitarString(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.GetGuitarString(mockContext)
}

func TestStringQueryController_GetGuitarString_NegativeUnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringQueryUseCase := mock_inputboundary.NewMockStringQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockContext.EXPECT().Param(gomock.Any()).Return("mock_param").Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockStringQueryUseCase.EXPECT().GetGuitarString(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.GetGuitarString(mockContext)
}