package main

import (
	"log"
	"net/http"

	"github.com/xeenhl/rest-api-demo/backend/routers"

	"github.com/urfave/negroni"
)

func main() {

	r := routers.InitRouters()

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":8081", n))
}
