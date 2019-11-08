package rest

import (
	"net/http"
	"lib/persistence"
	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, databaseHandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databaseHandler)

	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Method("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Method("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Method("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}
