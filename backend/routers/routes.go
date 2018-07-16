package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xeenhl/spender/backend/env"
)

type Env struct {
	env.Env
}

var environment *Env

func InitRouters(e *env.Env) *mux.Router {
	environment = &Env{*e}
	router := mux.NewRouter()
	router = SetSpendRoutes(router)
	return router
}

func SetSpendRoutes(router *mux.Router) *mux.Router {
	router.Handle("/spends",
		negroni.New(negroni.HandlerFunc(environment.GetSpendsHandler))).Methods("GET")
	router.Handle("/spends/{id}",
		negroni.New(negroni.HandlerFunc(environment.GetSingleSpendHandler))).Methods("GET")
	router.Handle("/spends/{id}",
		negroni.New(negroni.HandlerFunc(environment.UpdateSingleSpendHandler))).Methods("POST")

	return router
}
