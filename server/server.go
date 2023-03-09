package server

import (
	"net/http"
	"todo/handler"
	"todo/middleware"
	"todo/repo"

	"github.com/gorilla/mux"
)

type MyServer struct {
	router *mux.Router
}

func NewServer(rp *repo.Repo) MyServer {
	router := mux.NewRouter()
	routing := handler.Routing(rp)
	for _, route := range routing {
		router.Handle(route.Path, middleware.CORS(route.Handler)).Methods(route.Methods...)
	}

	return MyServer{
		router: router,
	}
}

func (srv MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}
