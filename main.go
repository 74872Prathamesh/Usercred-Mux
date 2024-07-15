package main

import (
	"net/http"
	"os"
	"usercred/component/user/controller"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter().StrictSlash(true)
	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})

	router = router.PathPrefix("/api/vi/userapp").Subrouter()
	controller1 := controller.NewUserController()
	controller1.RegisterRoutes(router)

	http.ListenAndServe(":4002", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
