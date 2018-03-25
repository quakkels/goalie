package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes routes ಠ益ಠ
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	return router
}
