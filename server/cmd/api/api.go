package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/felipejazz/ecommerce_go/cmd/service/user"
	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/server/v1").Subrouter()

	log.Println("Listening on", s.addr)
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}
