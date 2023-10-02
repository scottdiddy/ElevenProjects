package main

import (
	"moviesProject/src/api/router"
	"net/http"
)

func main() {
	router := router.InitializeRouter()

	http.Handle("/", router)
	println("server started on port 8080")
	http.ListenAndServe(":8080", nil)

}
