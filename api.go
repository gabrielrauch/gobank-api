package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string

	storage Storage
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}

func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		storage:       store,
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandler(s.handleAccount))

	log.Println("Iniciando servidor na porta:", s.listenAddress)

	return http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetAccount(w, r)
	} else if r.Method == http.MethodPost {
		return s.handleCreateAccount(w, r)
	} else if r.Method == http.MethodDelete {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := newAccount("John", "Doe")

	return returnJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// As helpers ficam no final do arquivo
func returnJSON(w http.ResponseWriter, status int, value any) error {
	// O Set precisa ser chamado antes do WriteHeader
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			// TODO: Precisa retornar os erros de forma correta de acordo com o status
			returnJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
