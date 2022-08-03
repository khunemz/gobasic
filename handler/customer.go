package handler

import (
	"encoding/json"
	"fmt"
	"go-hexagonal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	customerService service.CustomerService
}

// acts live constructor in OOP
func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{customerService: custSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		// no log because business layer has already logged
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	customer, err := h.customerService.GetCustomer(id)
	if err != nil {
		// no log because business layer has already logged
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
