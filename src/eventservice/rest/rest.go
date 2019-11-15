package rest

import (
	"net/http"
	"eventservice/lib/persistence"
	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, databaseHandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databaseHandler)

	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}
