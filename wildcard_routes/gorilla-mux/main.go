package main

import (
	testrunner "github.com/AgrimPrasad/go-mux-tests/wildcard_routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{path:js.*}/{blah}", testrunner.JSHandler1).Methods("GET")
	r.HandleFunc("/{path:js.*}/lib/{blah}", testrunner.JSHandler2).Methods("GET")       //fail (Got JSHandler1)
	r.HandleFunc("/{path:js.*}/lib/react/{blah}", testrunner.JSHandler3).Methods("GET") //fail (Got JSHandler1)
	r.HandleFunc("/some/path/1/something", testrunner.Handler1).Methods("GET")
	r.HandleFunc("/some/path/2/something", testrunner.Handler2).Methods("GET")
	r.HandleFunc("/teachers/list", testrunner.Handler3).Methods("GET")
	r.HandleFunc("/teachers/{id:\\d+}/profile", testrunner.Handler4).Methods("GET")
	r.HandleFunc("/teachers/{id:\\d+{.*}}/profile", testrunner.Handler5).Methods("GET")
	r.HandleFunc("/blah_lists/{listId:\\d+}", testrunner.ListHandlerGet).Methods("GET")
	r.HandleFunc("/blah_lists/{listId:\\d+}", testrunner.ListHandlerPost).Methods("POST")

	testrunner.TestWildcardRoutes(r)
}
