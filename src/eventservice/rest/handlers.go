package rest

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"eventservice/lib/persistence"
)

type eventServiceHandler struct {
	dbHandler persistence.DatabaseHandler
}

func NewEventHandler(databaseHandler persistence.DatabaseHandler) (*eventServiceHandler) {
	return &eventServiceHandler{
		dbHandler: databaseHandler,
	}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either search
						by id via /id/4
						by name via /name/coldplayconcert"}`)
	}

	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either search  
                                         	by id via /id/4
                                         	by name via /name/coldplayconcert"}`)
	}

	var event persistence.Event
	var err error
	switch(strings.ToLower(criteria)) {
	case "name":
		event, err = eh.dbHandler.FindEventByName(searchKey)
	case "id":
		id, err := hex.DecodeString(searchKey)
		if err == nil {
			event, err = eh.dbHandler.FindEvent(id)
		}
	}

	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbHandler.FindAllAvailableEvents()

	if err !=nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to find all available events %s"}`, err)
		return
	}

	w.Header().Set("Content-type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while encoding events to JSON %s"}`, err)
	}

}

func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while decoding event data %s"}`, err)
		return
	}

	id, err := eh.dbHandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while persisting event %d %s"}`, id, err)
	}

	fmt.Fprintf(w, `{"id": "%d"}`, id)
}
