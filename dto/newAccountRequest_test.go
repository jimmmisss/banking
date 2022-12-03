package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_account_type_is_not_saving_checking(t *testing.T) {
	// Arrange
	request := NewAccountRequest{
		AccountType: "invalid amount type",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Account type should be checking or saving" {
		t.Error(" Invalid message while validate account type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating amount type")
	}
}

func Test_should_return_error_when_amount_less_5000(t *testing.T) {
	// Arrange
	request := NewAccountRequest{
		AccountType: SAVING,
		Amount:      4500,
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "To open a new account you need to deposit at least 5000" {
		t.Error("Invalid message while validating amount")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating amount")
	}
}
