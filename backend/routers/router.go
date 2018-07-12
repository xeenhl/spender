package routers

import "github.com/gorilla/mux"

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = SetSpendRoutes(router)
	return router
}
