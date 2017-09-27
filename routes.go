package main

import (
	"github.com/infinityworks/go-common/router"
)

func createRoutes(c Controller) router.Routes {

	return router.Routes{

		router.Route{
			Name:        "FetchCountry",
			Method:      "GET",
			Pattern:     "/country",
			HandlerFunc: c.FetchCountry,
		},
	}
}
