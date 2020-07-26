package main

import (
	"net/http"

	"demosdata.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
