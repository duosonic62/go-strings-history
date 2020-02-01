package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/duosonic62/go-strings-history/internal/domain/service/mock_service"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary/mock_outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestStringQueryUseCase_GetGuitarString_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringQueryPresenter := mock_outputboundary.NewMockStringQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringQueryRepository := mock_repository.NewMockStringQueryRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(user(), nil).Times(1)
	mockStringQueryRepository.EXPECT().Find(gomock.Any()).Return(guitarStringOutput(), nil).Times(1)
	mockStringQueryPresenter.EXPECT().OutputGuitarString(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringQueryUseCase(
		mockAuthorizedService,
		mockStringQueryPresenter,
		mockErrorPresenter,
		mockStringQueryRepository,
	)
	useCase.GetGuitarString("mock_id", token(), mockContext)
}

func TestStringQueryUseCase_GetGuitarString_NegativeUnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringQueryPresenter := mock_outputboundary.NewMockStringQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringQueryRepository := mock_repository.NewMockStringQueryRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	mockStringQueryRepository.EXPECT().Find(gomock.Any()).Return(guitarStringOutput(), nil).Times(0)
	mockStringQueryPresenter.EXPECT().OutputGuitarString(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringQueryUseCase(
		mockAuthorizedService,
		mockStringQueryPresenter,
		mockErrorPresenter,
		mockStringQueryRepository,
	)
	useCase.GetGuitarString("mock_id", token(), mockContext)
}

func TestStringQueryUseCase_GetGuitarString_NegativeBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringQueryPresenter := mock_outputboundary.NewMockStringQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringQueryRepository := mock_repository.NewMockStringQueryRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(user(), nil).Times(1)
	mockStringQueryRepository.EXPECT().Find(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	mockStringQueryPresenter.EXPECT().OutputGuitarString(gomock.Any(), gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringQueryUseCase(
		mockAuthorizedService,
		mockStringQueryPresenter,
		mockErrorPresenter,
		mockStringQueryRepository,
	)
	useCase.GetGuitarString("mock_id", token(), mockContext)
}

// テストヘルパーメソッド
func guitarStringOutput() *output.GuitarStringOutput {
	return &output.GuitarStringOutput{
		ID:          "mock_ID",
		Name:        "mock_name",
		Description: "mock_description",
		Maker:       "mock_maker",
		ThinGauge:   9,
		ThickGauge:  42,
		Url:         "mock_url",
	}
}
