package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xeenhl/spender/backend/env"
	"github.com/xeenhl/spender/backend/models"
	"github.com/xeenhl/spender/backend/routers"

	"github.com/urfave/negroni"
)

func main() {

	db, err := models.InitDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	env := &env.Env{db}

	r := routers.InitRouters(env)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":8081", n))
}
