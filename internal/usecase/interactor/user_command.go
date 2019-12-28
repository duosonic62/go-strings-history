package interactor

import (
	"github.com/duosonic62/go-strings-history/internal/domain/factory"
	"github.com/duosonic62/go-strings-history/internal/domain/repository"
	"github.com/duosonic62/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic62/go-strings-history/internal/usecase/outputboundary"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input"
	"github.com/duosonic62/go-strings-history/pkg/usecase/input/command"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

type UserUseCaseInteractor struct {
	presenter   outputboundary.UserCommandPresenter
	repository  repository.UserRepository
	userFactory factory.UserFactory
}

// コンストラクタ
func NewUserCommandUseCase(
	presenter outputboundary.UserCommandPresenter,
	repository repository.UserRepository,
	userFactory factory.UserFactory,
) inputboundary.UserCommandUseCase {
	return UserUseCaseInteractor{
		presenter:   presenter,
		repository:  repository,
		userFactory: userFactory,
	}
}

func (useCase UserUseCaseInteractor) AddUser(data command.UserAddInputData, ctx input.Context) {
	// Userエンティティを作成
	user, err := useCase.userFactory.NewUser(data.Name, data.UID)
	if err != nil {
		// FIXME エラーハンドリング
		panic(err)
	}

	// エンティティを保存
	useCase.repository.Save(user)

	// presenterに表示処理を渡す
	outputData := output.UserAddOutputData{
		CreatedToken: user.Token,
	}
	useCase.presenter.OutputAddUser(outputData, ctx)
}
