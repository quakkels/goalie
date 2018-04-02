package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/quakkels/goalie/authentication"
	"github.com/quakkels/goalie/controllers"
)

// SetHelloRoutes will set hello routes ಠ益ಠ
func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle(
		"/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController))).Methods("GET")

	router.Handle(
		"/test/helloprotected",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloProtected))).Methods("GET")

	return router
}
