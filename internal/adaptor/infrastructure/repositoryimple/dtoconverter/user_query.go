package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/duosonic62/go-strings-history/pkg/usecase/output"
)

func ConvertUser(member *models.Member) output.UserOutput {
	return output.UserOutput{
		Name: member.Name,
	}
}
