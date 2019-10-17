package router

import (
	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
	Routes "github.com/projects/vue-go-1st/api/router/routes"
	v1Routes "github.com/projects/vue-go-1st/api/router/routes/v1"
	"net/http"
)

type RouteHandler struct {
	Router *mux.Router
}

const (
	static = "/static/"
)

func (r *RouteHandler) AttachSubRouterWithMiddleware(
	path string,
	subRoutes Routes.Routes,
	Middleware mux.MiddlewareFunc,
) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(Middleware)

	for _, route := range subRoutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return SubRouter
}

func NewRouter(db *xorm.Engine) *RouteHandler {
	var router RouteHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix("/get").
		Handler(http.StripPrefix("/get", http.FileServer(http.Dir("."+static))))

	router.Router.Use(Routes.Middleware)

	routes := Routes.GetRoutes(db)

	for _, route := range routes {
		router.Router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	V1Routes := v1Routes.GetRoutes(db)
	for name, pack := range V1Routes {
		router.AttachSubRouterWithMiddleware(
			name,
			pack.Routes,
			pack.Middleware,
		)

	}

	return &router
}
