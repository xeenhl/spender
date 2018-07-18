package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xeenhl/spender/backend/models"
)

func (env *Env) GetSpendsHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	spends, err := env.DB.GetAllSpends(r.Context())
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
	}
	json.NewEncoder(rw).Encode(spends)
}

func (env *Env) GetSingleSpendHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	props := mux.Vars(r)
	id, err := strconv.Atoi(props["id"])

	if props == nil || err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Bad Request!"))
		return
	}

	if spend, err := env.DB.GetSpendById(id, r.Context()); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Not Found!"))
	} else {
		json.NewEncoder(rw).Encode(spend)
	}
}

func (env *Env) UpdateSingleSpendHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	props := mux.Vars(r)
	id, err := strconv.Atoi(props["id"])
	if props == nil || err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Not Found!"))
		return
	}

	var newData models.Spend
	decodeError := json.NewDecoder(r.Body).Decode(&newData)

	if decodeError != nil {
		rw.Write([]byte(decodeError.Error()))
		return
	}

	spend, err := env.DB.UpdateSpend(id, newData, r.Context())

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
	}

	json.NewEncoder(rw).Encode(spend)
}
