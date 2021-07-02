package main

import (
	"net/http"

	"github.com/cameronperera/web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
