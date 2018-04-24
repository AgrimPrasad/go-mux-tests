package wildcard_routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func TestWildcardRoutes(r http.Handler) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/js/1.js", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/js/lib/2.js", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/js/lib/react/3.js", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/js%2Flib/3.js", nil)
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/some/path/1/something", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/some/path/2/something", nil)
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/teachers/list", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/teachers/2/profile", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/teachers/2{blah}/profile", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("POST", "/teachers/2/profile", nil) //ignored
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/blah_lists/3", nil)
	r.ServeHTTP(w, req)
	req, _ = http.NewRequest("POST", "/blah_lists/3", nil)
	r.ServeHTTP(w, req)
}

func JSHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("JSHandler1", r.URL.Path)
}

func JSHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("JSHandler2", r.URL.Path)
}

func JSHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("JSHandler3", r.URL.Path)
}

func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler1", r.URL.Path)
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler2", r.URL.Path)
}

func Handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler3", r.URL.Path)
}

func Handler4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler4", r.URL.Path)
}

func Handler5(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler5", r.URL.Path)
}

func ListHandlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listHandlerGet", r.URL.Path)
}

func ListHandlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listHandlerPost", r.URL.Path)
}
