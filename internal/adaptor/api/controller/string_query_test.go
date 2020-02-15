package controller

import (
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary/mock_inputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/query"
	"github.com/golang/mock/gomock"
	"github.com/volatiletech/null"
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
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
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

	mockContext.EXPECT().Param(gomock.Any()).Return("mock_param").Times(0)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockStringQueryUseCase.EXPECT().GetGuitarString(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.GetGuitarString(mockContext)
}

func TestStringQueryController_SearchGuitarString_PositiveQueries(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringQueryUseCase := mock_inputboundary.NewMockStringQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	expectQuery := query.SearchGuitarString{
		Name: null.StringFrom("mock_name"),
		Maker: null.StringFrom("mock_maker"),
		ThinGauge: null.NewInt(10, true),
		ThickGauge: null.NewInt(46, true),
	}

	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(0)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
	// クエリ全てあり
	mockContext.EXPECT().GetQuery("name").Return("mock_name", true).Times(1)
	mockContext.EXPECT().GetQuery("maker").Return("mock_maker", true).Times(1)
	mockContext.EXPECT().GetQuery("thinGauge").Return("10", true).Times(1)
	mockContext.EXPECT().GetQuery("thickGauge").Return("46", true).Times(1)
	mockStringQueryUseCase.EXPECT().SearchGuitarString(expectQuery, gomock.Any(), gomock.Any())

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.SearchGuitarString(mockContext)
}

func TestStringQueryController_SearchGuitarString_PositiveNoQueries(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringQueryUseCase := mock_inputboundary.NewMockStringQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	expectQuery := query.SearchGuitarString{
		Name: null.NewString("", false),
		Maker: null.NewString("", false),
		ThinGauge: null.NewInt(0, false),
		ThickGauge: null.NewInt(0, false),
	}

	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(0)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
	mockContext.EXPECT().GetQuery("name").Return("", false).Times(1)
	mockContext.EXPECT().GetQuery("maker").Return("", false).Times(1)
	mockContext.EXPECT().GetQuery("thinGauge").Return("", false).Times(1)
	mockContext.EXPECT().GetQuery("thickGauge").Return("", false).Times(1)
	mockStringQueryUseCase.EXPECT().SearchGuitarString(expectQuery, gomock.Any(), gomock.Any())

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.SearchGuitarString(mockContext)
}

func TestStringQueryController_SearchGuitarString_NegativeBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringQueryUseCase := mock_inputboundary.NewMockStringQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(1)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(0)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("mock_token").Times(1)
	mockContext.EXPECT().GetQuery("name").Return("", false).Times(1)
	mockContext.EXPECT().GetQuery("maker").Return("", false).Times(1)
	mockContext.EXPECT().GetQuery("thinGauge").Return("invalid_value", true).Times(1)
	mockContext.EXPECT().GetQuery("thickGauge").Return("invalid_value", false).Times(0)
	mockStringQueryUseCase.EXPECT().SearchGuitarString(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.SearchGuitarString(mockContext)
}

func TestStringQueryController_SearchGuitarString_NegativeUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStringQueryUseCase := mock_inputboundary.NewMockStringQueryUseCase(ctrl)
	mockErrorUseCase := mock_inputboundary.NewMockErrorUseCase(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockErrorUseCase.EXPECT().BadRequestError(gomock.Any(), gomock.Any()).Times(0)
	mockErrorUseCase.EXPECT().UnauthorizedError(gomock.Any(), gomock.Any()).Times(1)
	mockContext.EXPECT().GetHeader(gomock.Any()).Return("").Times(1)
	mockContext.EXPECT().GetQuery("name").Return("", false).Times(0)
	mockContext.EXPECT().GetQuery("maker").Return("", false).Times(0)
	mockContext.EXPECT().GetQuery("thinGauge").Return("", false).Times(0)
	mockContext.EXPECT().GetQuery("thickGauge").Return("", false).Times(0)
	mockStringQueryUseCase.EXPECT().SearchGuitarString(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	controller := NewStringQueryController(mockStringQueryUseCase, mockErrorUseCase)
	controller.SearchGuitarString(mockContext)
}