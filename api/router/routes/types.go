package routes

import (
	"net/http"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
