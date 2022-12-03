package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jimmmisss/banking-lib/logger"
	"github.com/jimmmisss/banking/domain"
	"github.com/jimmmisss/banking/service"
	"github.com/jimmmisss/banking/util"
	"log"
	"net/http"
)

func Start() {
	env := util.EnvCheck()
	router := mux.NewRouter()

	dbClient := util.EnvDB()
	customerRepository := domain.NewCustomerRepositoryDB(dbClient)
	accountRepository := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepository)}
	ah := AccountHandler{service.NewAccountService(accountRepository)}

	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getFindById).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	address := env.Server.Address
	port := env.Server.Port
	logger.Info(fmt.Sprintf("Start server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
