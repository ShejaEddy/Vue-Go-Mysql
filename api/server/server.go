package server

import (
	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
	routerFactory "github.com/projects/vue-go-1st/api/router"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	Port       int
	Addr       string
	HTTPServer *http.Server
}

func (s *Server) Start() {
	log.Println("Server has started on port", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())
}

func NewServer(port int, db *xorm.Engine) *Server {
	var server Server

	server.Port = port
	server.Addr = ":" + strconv.Itoa(port)

	router := routerFactory.NewRouter(db)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "X-Access-Token", "Cashe-Control"}),
		handlers.ExposedHeaders([]string{}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(router.Router))

	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	server.HTTPServer = &http.Server{
		Addr:           server.Addr,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server
}
