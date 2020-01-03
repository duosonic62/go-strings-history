package interactor

import (
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/domain/valueobject"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserUseCaseInteractor struct {
	presenter      outputboundary.UserCommandPresenter
	errorPresenter outputboundary.ErrorPresenter
	repository     repository.UserCommandRepository
	userFactory    factory.UserFactory
}

// コンストラクタ
func NewUserCommandUseCase(
	presenter outputboundary.UserCommandPresenter,
	errorPresenter outputboundary.ErrorPresenter,
	repository repository.UserCommandRepository,
	userFactory factory.UserFactory,
) inputboundary.UserCommandUseCase {
	return UserUseCaseInteractor{
		presenter:      presenter,
		errorPresenter: errorPresenter,
		repository:     repository,
		userFactory:    userFactory,
	}
}

func (useCase UserUseCaseInteractor) Add(data command.UserAddInputData, ctx input.Context) {
	// Userエンティティを作成
	user, err := useCase.userFactory.NewUser(data.Name, data.UID)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	// エンティティを保存
	err = useCase.repository.Save(user)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	// presenterに表示処理を渡す
	outputData := output.UserAddOutputData{
		CreatedToken: user.GetToken().GetToken(),
	}
	useCase.presenter.OutputAddUser(outputData, ctx)
}

func (useCase UserUseCaseInteractor) Edit(token valueobject.AuthorizationToken, data command.UserEditInputData, ctx input.Context) {
	// Userエンティティを作成
	user, err := useCase.userFactory.Find(token)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	// エンティティを保存
	err = useCase.repository.Save(user)
	if err != nil {
		useCase.errorPresenter.OutputError(ctx, err)
		return
	}

	// presenterに表示処理を渡す
	outputData := output.UserAddOutputData{
		CreatedToken: user.GetToken().GetToken(),
	}
	useCase.presenter.OutputAddUser(outputData, ctx)
}
