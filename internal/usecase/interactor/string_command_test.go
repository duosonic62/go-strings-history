package interactor

import (
	"errors"
	"github.com/duosonic62/go-strings-history/internal/domain/entity"
	"github.com/duosonic62/go-strings-history/internal/domain/factory/mock_factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository/mock_repository"
	"github.com/duosonic62/go-strings-history/internal/domain/service/mock_service"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary/mock_outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/mock_input"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestStringCommandUseCaseImpl_Add_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Times(1).Return(user(), nil)
	mockStringFactory.EXPECT().
		NewString(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(1).
		Return(guitarString(), nil)
	mockStringCommandRepository.EXPECT().Save(gomock.Any()).Times(1).Return(nil)
	mockStringCommandPresenter.EXPECT().OutputAddString(gomock.Any()).Times(1)
	mockContext := mock_input.NewMockContext(ctrl)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Add(stringRegisterInputData(), token(), mockContext)
}

func TestStringCommandUseCaseImpl_Add_Negative_UnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Times(1).Return(nil, errors.New("error"))
	mockStringFactory.EXPECT().
		NewString(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)
	mockStringCommandRepository.EXPECT().Save(gomock.Any()).Times(0)
	mockStringCommandPresenter.EXPECT().OutputAddString(gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockContext := mock_input.NewMockContext(ctrl)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Add(stringRegisterInputData(), token(), mockContext)
}

func TestStringCommandUseCaseImpl_Add_Negative_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Times(1).Return(user(), nil)
	mockStringFactory.EXPECT().
		NewString(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil, errors.New("error"))
	mockStringCommandRepository.EXPECT().Save(gomock.Any()).Times(0)
	mockStringCommandPresenter.EXPECT().OutputAddString(gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockContext := mock_input.NewMockContext(ctrl)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Add(stringRegisterInputData(), token(), mockContext)
}

func TestStringCommandUseCaseImpl_Add_Negative_DbError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Times(1).Return(user(), nil)
	mockStringFactory.EXPECT().
		NewString(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(1).
		Return(guitarString(), nil)
	mockStringCommandRepository.EXPECT().Save(gomock.Any()).Times(1).Return(errors.New("error"))
	mockStringCommandPresenter.EXPECT().OutputAddString(gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)
	mockContext := mock_input.NewMockContext(ctrl)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Add(stringRegisterInputData(), token(), mockContext)
}

func TestStringCommandUseCaseImpl_Update_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(nil, nil).Times(1)
	mockStringCommandRepository.EXPECT().FindByID(gomock.Any()).Return(guitarString(), nil).Times(1)
	mockStringCommandRepository.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
	mockStringCommandPresenter.EXPECT().OutputUpdateString(gomock.Any()).Times(1)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Update("mock_id", stringRegisterInputData(), token(), mockContext)
}

func TestStringCommandUseCaseImpl_Update_Negative_UnAuthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	mockStringCommandRepository.EXPECT().FindByID(gomock.Any()).Return(guitarString(), nil).Times(0)
	mockStringCommandRepository.EXPECT().Update(gomock.Any()).Return(nil).Times(0)
	mockStringCommandPresenter.EXPECT().OutputUpdateString(gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Update("mock_id", stringRegisterInputData(), token(), mockContext)
}

func TestStringCommandUseCaseImpl_Update_Negative_BadRequest_InvalidId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(nil, nil).Times(1)
	mockStringCommandRepository.EXPECT().FindByID(gomock.Any()).Return(nil, errors.New("error")).Times(1)
	mockStringCommandRepository.EXPECT().Update(gomock.Any()).Return(nil).Times(0)
	mockStringCommandPresenter.EXPECT().OutputUpdateString(gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Update("mock_id", stringRegisterInputData(), token(), mockContext)
}


func TestStringCommandUseCaseImpl_Update_Negative_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorizedService := mock_service.NewMockAuthorizationService(ctrl)
	mockStringCommandPresenter := mock_outputboundary.NewMockStringCommandPresenter(ctrl)
	mockErrorPresenter := mock_outputboundary.NewMockErrorPresenter(ctrl)
	mockStringFactory := mock_factory.NewMockStringFactory(ctrl)
	mockStringCommandRepository := mock_repository.NewMockStringCommandRepository(ctrl)
	mockContext := mock_input.NewMockContext(ctrl)

	mockAuthorizedService.EXPECT().Authorized(gomock.Any()).Return(nil, nil).Times(1)
	mockStringCommandRepository.EXPECT().FindByID(gomock.Any()).Return(guitarString(), nil).Times(1)
	mockStringCommandRepository.EXPECT().Update(gomock.Any()).Return(errors.New("error")).Times(1)
	mockStringCommandPresenter.EXPECT().OutputUpdateString(gomock.Any()).Times(0)
	mockErrorPresenter.EXPECT().OutputError(gomock.Any(), gomock.Any()).Times(1)

	useCase := NewStringCommandUseCase(
		mockAuthorizedService,
		mockStringCommandPresenter,
		mockErrorPresenter,
		mockStringFactory,
		mockStringCommandRepository,
	)
	useCase.Update("mock_id", stringRegisterInputData(), token(), mockContext)
}

// テストヘルパーメソッド
func guitarString() *entity.GuitarString {
	gauge, _ := valueobject.NewStringGauge(9, 42)
	guitarString, _ := entity.NewGuitarString(
		"mock_id",
		"mock_name",
		"mock_description",
		"mock_maker",
		gauge,
		"mock_url",
	)
	return guitarString
}

func updatedGuitarString() *entity.GuitarString {
	gauge, _ := valueobject.NewStringGauge(10, 46)
	guitarString, _ := entity.NewGuitarString(
		"mock_id",
		"new_mock_name",
		"new_mock_description",
		"new_mock_maker",
		gauge,
		"new_mock_url",
	)
	return guitarString
}

func stringRegisterInputData() command.StringRegisterInputData {
	return command.StringRegisterInputData{
		Name:        "new_mock_name",
		Description: "new_mock_description",
		Maker:       "new_mock_description",
		ThinGauge:   10,
		ThickGauge:  46,
		Url:         "new_mock_url",
	}
}
