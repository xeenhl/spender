package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xeenhl/spender/backend/models"

	"github.com/urfave/negroni"
)

type Env struct {
	db models.Datastore
}

var env *Env

func main() {

	db, err := models.InitDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	env = &Env{db}

	r := InitRouters()

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":8081", n))
}

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = SetSpendRoutes(router)
	return router
}

func SetSpendRoutes(router *mux.Router) *mux.Router {
	router.Handle("/spends",
		negroni.New(negroni.HandlerFunc(env.GetSpendsHandler))).Methods("GET")
	router.Handle("/spends/{id}",
		negroni.New(negroni.HandlerFunc(env.GetSingleSpendHandler))).Methods("GET")
	router.Handle("/spends/{id}",
		negroni.New(negroni.HandlerFunc(env.UpdateSingleSpendHandler))).Methods("POST")

	return router
}

func (env *Env) GetSpendsHandler(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	spends, err := env.db.GetAllSpends()
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

	if spend, err := env.db.GetSpendById(id); err != nil {
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

	spends, err := env.db.GetAllSpends()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
	}

	var newData models.Spend
	decodeError := json.NewDecoder(r.Body).Decode(&newData)

	if decodeError != nil {
		rw.Write([]byte(decodeError.Error()))
		return
	}

	for _, s := range spends {
		if s.ID == id {
			s.Update(newData)
			json.NewEncoder(rw).Encode(s)
			return
		}
	}

	rw.Write([]byte("Not Found!"))
}
