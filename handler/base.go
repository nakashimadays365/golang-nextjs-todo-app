package handler

import (
	"net/http"
	"todo/repo"
)

type Handler struct {
	rp    *repo.Repo
	files []string
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Methods []string
}

func Routing(rp *repo.Repo) []Route {

	h := Handler{
		rp: rp,
		files: []string{
			"tpls/index.html",
		},
	}
	return []Route{
		{
			Path:    "/",
			Handler: h.Index,
			Methods: []string{http.MethodGet, http.MethodOptions},
		},
		{
			Path:    "/create",
			Handler: h.Create,
			Methods: []string{http.MethodPost, http.MethodOptions},
		},
		{
			Path:    "/update",
			Handler: h.Update,
			Methods: []string{http.MethodPost, http.MethodOptions},
		},
		{
			Path:    "/del",
			Handler: h.Delete,
			Methods: []string{http.MethodPost, http.MethodOptions},
		},
	}
}
