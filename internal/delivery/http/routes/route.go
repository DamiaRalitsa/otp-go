package route

import (
	"net/http"

	handler "bookingtogo/internal/delivery/http"

	"github.com/gorilla/mux"
)

func NewRouter(
	customerHandler *handler.CustomerHandler,
	familyHandler *handler.FamilyHandler,
	nationalityHandler *handler.NationalityHandler,
) *mux.Router {
	r := mux.NewRouter()

	// Nationality routes
	r.HandleFunc("/nationalities", nationalityHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/nationalities/{id:[0-9]+}", nationalityHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/nationalities", nationalityHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/nationalities/{id:[0-9]+}", nationalityHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/nationalities/{id:[0-9]+}", nationalityHandler.Delete).Methods(http.MethodDelete)

	// Customer routes
	r.HandleFunc("/customers", customerHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/customers/{id:[0-9]+}", customerHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customers", customerHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id:[0-9]+}", customerHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/customers/{id:[0-9]+}", customerHandler.Delete).Methods(http.MethodDelete)

	// Family routes
	r.HandleFunc("/customers/{cst_id:[0-9]+}/family", familyHandler.GetAllByCustomer).Methods(http.MethodGet)
	r.HandleFunc("/family/{ft_id:[0-9]+}", familyHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customers/{cst_id:[0-9]+}/family", familyHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/family/{ft_id:[0-9]+}", familyHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/family/{ft_id:[0-9]+}", familyHandler.Delete).Methods(http.MethodDelete)

	return r
}
