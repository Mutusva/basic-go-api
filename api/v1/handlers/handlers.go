package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api/api"
	"net/http"
	"strconv"
)

func Demofunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["type"]

	if t == "" {
		writeJsonResponse(w, http.StatusBadRequest, api.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "type cannot be empty",
		})
		return
	}
	tp, err := strconv.Atoi(t)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, api.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: "type cannot be empty",
		})
		return
	}
	var demos []api.Demo
	for _, d := range api.Demos {
		if d.Type == tp {
			demos = append(demos, d)
		}
	}

	writeJsonResponse(w, http.StatusOK, demos)
	return
}

func writeJsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
