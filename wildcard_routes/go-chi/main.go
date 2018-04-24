package main

import (
	testrunner "github.com/AgrimPrasad/go-mux-tests/wildcard_routes"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{path:js.*}/*", testrunner.JSHandler1)
	r.Get("/{path:js.*}/lib/*", testrunner.JSHandler2)
	r.Get("/{path:js.*}/lib/react/*", testrunner.JSHandler3)
	r.Get("/some/path/1/something", testrunner.Handler1)
	r.Get("/some/path/2/something", testrunner.Handler2)
	r.Get("/teachers/list", testrunner.Handler3)
	r.Get("/teachers/{id:\\d+}/profile", testrunner.Handler4)
	r.Get("/teachers/{id:\\d+{.*}}/profile", testrunner.Handler5)
	r.Get("/blah_lists/{listId:\\d+}", testrunner.ListHandlerGet)
	r.Post("/blah_lists/{listId:\\d+}", testrunner.ListHandlerPost)

	testrunner.TestWildcardRoutes(r)
}
