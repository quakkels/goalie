package main

import (
	"net/http"

	"github.com/codegangsta/negroni"

	"github.com/quakkels/goalie/routers"
	"github.com/quakkels/goalie/settings"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
