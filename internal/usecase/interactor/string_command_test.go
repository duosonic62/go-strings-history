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
		NewString(gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any()).
		Times(1).
		Return(guitarString(),nil)
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
	useCase.Add(stringAddInputData(), token(), mockContext)
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
		NewString(gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any()).
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
	useCase.Add(stringAddInputData(), token(), mockContext)
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
		NewString(gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any()).
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
	useCase.Add(stringAddInputData(), token(), mockContext)
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
		NewString(gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any(),gomock.Any()).
		Times(1).
		Return(guitarString(),nil)
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
	useCase.Add(stringAddInputData(), token(), mockContext)
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

func stringAddInputData() command.StringRegisterInputData {
	return command.StringRegisterInputData{}
}