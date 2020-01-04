package dtoconverter

import (
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/models"
	"github.com/volatiletech/null"
	"testing"
	"time"
)

func TestToUserEntity_Positive(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	dummyModel := models.Member{
		ID:           "mock_id",
		UID:          "mock_uid",
		Name:         "mock_name",
		Token:        "mock_token",
		TokenExpired: time.Date(2019, 10, 10, 10, 0, 0, 0, jst),
		IsDeleted:    false,
		CreatedAt:    time.Date(2019, 9, 10, 10, 0, 0, 0, jst),
		UpdatedAt:    null.TimeFrom(time.Date(2019, 9, 10, 10, 0, 0, 0, jst)),
	}
	actual, err := ToUserEntity(&dummyModel)
	if err != nil {
		t.Error("Expected: error is nil")
	}

	if actual.GetName() != "mock_name" {
		t.Error("Expected mock_name but Actual " + actual.GetName())
	}
	if actual.GetID() != "mock_id" {
		t.Error("Expected mock_name but Actual " + actual.GetID())
	}
	if actual.GetUID() != "mock_uid" {
		t.Error("Expected mock_uid but Actual " + actual.GetUID())
	}
	if actual.GetToken().GetToken() != "mock_token" {
		t.Error("Expected mock_token but Actual " + actual.GetToken().GetToken())
	}
}

func TestToUserEntity_Negative_InvalidToken(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	dummyModel := models.Member{
		ID:           "mock_id",
		UID:          "mock_uid",
		Name:         "mock_name",
		Token:        "", // invalid token
		TokenExpired: time.Date(2019, 10, 10, 10, 0, 0, 0, jst),
		IsDeleted:    false,
		CreatedAt:    time.Date(2019, 9, 10, 10, 0, 0, 0, jst),
		UpdatedAt:    null.TimeFrom(time.Date(2019, 9, 10, 10, 0, 0, 0, jst)),
	}

	_, err := ToUserEntity(&dummyModel)
	if err == nil {
		t.Error("Expected: error is not nil")
	}
}

func TestToUserEntity_Negative_InvalidUserEntity(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	dummyModel := models.Member{
		ID:           "", // invalid id
		UID:          "mock_uid",
		Name:         "mock_name",
		Token:        "mock_token",
		TokenExpired: time.Date(2019, 10, 10, 10, 0, 0, 0, jst),
		IsDeleted:    false,
		CreatedAt:    time.Date(2019, 9, 10, 10, 0, 0, 0, jst),
		UpdatedAt:    null.TimeFrom(time.Date(2019, 9, 10, 10, 0, 0, 0, jst)),
	}

	_, err := ToUserEntity(&dummyModel)
	if err == nil {
		t.Error("Expected: error is not nil")
	}
}
