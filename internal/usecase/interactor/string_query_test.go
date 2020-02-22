package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/duosonic62/go-strings-history/internal/domain/service/mock_service"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary/mock_outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/query"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
	"github.com/golang/mock/gomock"
	"github.com/volatiletech/null"
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

func TestStringQueryUseCase_SearchGuitarString_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringQueryPresenter := mock_outputboundary.NewMockStringQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringQueryRepository := mock_repository.NewMockStringQueryRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(user(), nil).Times(1)
	mockStringQueryRepository.EXPECT().
		Search(null.NewString("mock_name", true), null.NewString("mock_maker", true), null.NewInt(10, true), null.NewInt(46, true)).
		Return(guitarStringOutputs(), nil).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(0)
	mockStringQueryPresenter.EXPECT().OutputGuitarStrings(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringQueryUseCase(
		mockAuthorizedService,
		mockStringQueryPresenter,
		mockErrorPresenter,
		mockStringQueryRepository,
	)
	useCase.SearchGuitarString(guitarStringSearchQUery(), token(), mockContext)
}

func TestStringQueryUseCase_SearchGuitarString_NegativeUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringQueryPresenter := mock_outputboundary.NewMockStringQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringQueryRepository := mock_repository.NewMockStringQueryRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	mockStringQueryRepository.EXPECT().
		Search(null.NewString("mock_name", true), null.NewString("mock_maker", true), null.NewInt(10, true), null.NewInt(46, true)).
		Return(guitarStringOutputs(), nil).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockStringQueryPresenter.EXPECT().OutputGuitarStrings(gomock.Any(), gomock.Any()).Times(0)

	useCase := NewStringQueryUseCase(
		mockAuthorizedService,
		mockStringQueryPresenter,
		mockErrorPresenter,
		mockStringQueryRepository,
	)
	useCase.SearchGuitarString(guitarStringSearchQUery(), token(), mockContext)
}

func TestStringQueryUseCase_SearchGuitarString_NegativeDBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringQueryPresenter := mock_outputboundary.NewMockStringQueryPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringQueryRepository := mock_repository.NewMockStringQueryRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(user(), nil).Times(1)
	mockStringQueryRepository.EXPECT().
		Search(null.NewString("mock_name", true), null.NewString("mock_maker", true), null.NewInt(10, true), null.NewInt(46, true)).
		Return(nil, errors.New("error")).Times(1)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockStringQueryPresenter.EXPECT().OutputGuitarStrings(gomock.Any(), gomock.Any()).Times(0)

	useCase := NewStringQueryUseCase(
		mockAuthorizedService,
		mockStringQueryPresenter,
		mockErrorPresenter,
		mockStringQueryRepository,
	)
	useCase.SearchGuitarString(guitarStringSearchQUery(), token(), mockContext)
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

func guitarStringOutputs() *[]output.GuitarStringOutput {
	outputs := make([]output.GuitarStringOutput, 2)
	outputs[0] = *guitarStringOutput()
	outputs[1] = *guitarStringOutput()
	return &outputs
}

func guitarStringSearchQUery() query.SearchGuitarString {
	return query.SearchGuitarString{
		Name: null.NewString("mock_name", true),
		Maker: null.NewString("mock_maker", true),
		ThinGauge: null.NewInt(10, true),
		ThickGauge: null.NewInt(46, true),
	}
}
