package interactor

import (
	"github.com/duosonic/go-strings-history/internal/usecase/inputboundary"
	"github.com/duosonic/go-strings-history/pkg/usecase/input"
	"github.com/duosonic/go-strings-history/pkg/usecase/input/command"
	"github.com/duosonic/go-strings-history/pkg/usecase/output"
)

type UserUseCaseInteractor struct {
	presenter   outputboundary.UserPresenter
	repository  repository.UserRepository
	idFactory   factory.IDFactory
	userFactory factory.UserFactory
}

// コンストラクタ
func NewUserUsecase(
	presenter outputboundary.UserPresenter,
	repository repository.UserRepository,
	idFactory factory.IDFactory,
	userFactory factory.UserFactory,
) inputboundary.UserCommandUseCase {
	return UserUseCaseInteractor{
		presenter:   presenter,
		repository:  repository,
		idFactory:   idFactory,
		userFactory: userFactory,
	}
}

func (useCase UserUseCaseInteractor) AddUser(data command.UserAddInputData, ctx input.Context) {
	// Userエンティティを作成
	user, err := useCase.userFactory.NewUser(data.Name, data.Role)
	if err != nil {
		// FIXME エラーハンドリング
		panic(err)
	}

	// エンティティを保存
	useCase.repository.Save(user)

	// presenterに表示処理を渡す
	outputData := output.UserAddOutputData{
		CreatedId: user.ID,
	}
	useCase.presenter.OutputAddUser(outputData, ctx)
}
