package router

import (
	"moviesProject/src/api/controllers"
	"moviesProject/src/middlewares/authentication"
	"moviesProject/src/middlewares/validation"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()

	getRouter := apirouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/movies", controllers.GetMoviesController)
	getRouter.HandleFunc("/movie/{id}", controllers.GetMovieController)

	createRouter := apirouter.Methods(http.MethodPost).Subrouter()
	createRouter.HandleFunc("/movie", controllers.CreateMovieController)
	createRouter.Use(validation.ValidatorforCreateRequest)

	updateRouter := apirouter.Methods(http.MethodPatch).Subrouter()
	updateRouter.HandleFunc("/movie", controllers.UpdateMovieController)
	updateRouter.Use(authentication.MiddlewareAuthenticateUserToken, validation.ValidatorforUpdateRequest)

	deleteRouter := apirouter.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/movie", controllers.DeleteMovieController)
	deleteRouter.Use(authentication.MiddlewareAuthenticateUserToken)

	return router
}
