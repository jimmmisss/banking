package app

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking/dto"
	"github.com/jimmmisss/banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers
var mockCustomerService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCustomerService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockCustomerService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
		{Id: "0001", Name: "Wesley Pereira", City: "Palhoça/SC", Zipcode: "01", DateOfBirth: "1982-09-13", Status: "1"},
		{Id: "0002", Name: "Isadora Pereira", City: "Pindaré Mirim/MA", Zipcode: "01", DateOfBirth: "2013-05-04", Status: "1"},
	}
	mockCustomerService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockCustomerService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("same database error"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
