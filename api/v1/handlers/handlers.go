package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-api/datastore"
	"go-api/models"
	"net/http"
	"strconv"
)

type LocationApp struct {
	DS datastore.Datastore
}

func (app *LocationApp) NowLocation(w http.ResponseWriter, r *http.Request) {
	var req models.Location
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, "Wrong input")
		return
	}

	vars := mux.Vars(r)
	userId := vars["user_id"]
	if userId == "" {
		writeJsonResponse(w, http.StatusBadRequest, "Wrong input")
		return
	}

	//TODO: make validations work
	app.DS.UpdateLocation(userId, &req)
	writeJsonResponse(w, http.StatusOK, fmt.Sprintf("%s locations updates", userId))

	return
}

func (app *LocationApp) GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	maxReq := vars["max"]

	max, err := strconv.Atoi(maxReq)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, "Wrong input")
		return
	}

	locs, err := app.DS.GetLocation(userId, max)
	writeJsonResponse(w, http.StatusOK, locs)
	return
}

func (app *LocationApp) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	if userId == "" {
		writeJsonResponse(w, http.StatusBadRequest, "Wrong input")
		return
	}

	err := app.DS.Delete(userId)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, "Could not delete user location")
		return
	}

	writeJsonResponse(w, http.StatusOK, "User locations delete")
	return

}

func writeJsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
