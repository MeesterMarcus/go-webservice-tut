package main

import (
	"net/http"

	"github.com/meestermarcus/go-webservice-tut/controllers"
)

/*
	Golang sandbox
*/
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
