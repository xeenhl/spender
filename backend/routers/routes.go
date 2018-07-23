package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xeenhl/spender/backend/authentication"
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
	router = SetLoginRoutes(router)
	return router
}

func SetSpendRoutes(router *mux.Router) *mux.Router {
	router.Handle("/spends",
		negroni.New(
			negroni.HandlerFunc(authentication.AuhtWithToken),
			negroni.HandlerFunc(environment.GetSpendsHandler))).Methods("GET")
	router.Handle("/spends/{id}",
		negroni.New(
			negroni.HandlerFunc(authentication.AuhtWithToken),
			negroni.HandlerFunc(environment.GetSingleSpendHandler))).Methods("GET")
	router.Handle("/spends",
		negroni.New(
			negroni.HandlerFunc(authentication.AuhtWithToken),
			negroni.HandlerFunc(environment.AddSingleSpendHandler))).Methods("POST")

	return router
}

func SetLoginRoutes(router *mux.Router) *mux.Router {
	router.Handle("/login",
		negroni.New(
			negroni.HandlerFunc(environment.Login),
			negroni.HandlerFunc(authentication.SignWithNewToken))).Methods("POST")
	router.Handle("/sigin",
		negroni.New(
			negroni.HandlerFunc(environment.SignIn),
			negroni.HandlerFunc(authentication.SignWithNewToken))).Methods("POST")

	return router
}
