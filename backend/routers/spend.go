package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xeenhl/rest-api-demo/backend/controllers"
)

func SetSpendRoutes(router *mux.Router) *mux.Router {
	router.Handle("/spends",
		negroni.New(negroni.HandlerFunc(controllers.GetSpendsHandler))).Methods("GET")
	return router
}
