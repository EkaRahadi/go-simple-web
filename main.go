package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)
}

func QueryHandlerArray(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	var name []string = queries["name"]
	w.Header().Add("x-powered-by", "eka-rahadi")
	connection := r.Header.Get("connection")
	fmt.Println(connection)
	fmt.Fprintf(w, "Hello, %s %s", name[0], name[1])
}

func FormPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello, %s %s", firstName, lastName)
}

func CookieAndStatusCodeHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	if firstName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No First Name!")
	} else {
		cookie := new(http.Cookie)
		cookie.Name = "first_name"
		cookie.Value = firstName
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Hello, %s", firstName)
	}
}

func GetCookieAndStatusCodeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("first_name")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No Cookie!")
	} else {
		fmt.Fprintf(w, "Hello, %s", cookie.Value)
	}
}

func ServeFileEmbedHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("first_name") != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

//go:embed resources
var resources embed.FS

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/not_found.html
var resourceNotFound string

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/say-hello", QueryHandler)
	mux.HandleFunc("/say-hello-array", QueryHandlerArray)
	mux.HandleFunc("/say-hello-post", FormPostHandler)
	mux.HandleFunc("/say-hello-post-cookie", CookieAndStatusCodeHandler)
	mux.HandleFunc("/say-hello-get-cookie", GetCookieAndStatusCodeHandler)

	directory, _ := fs.Sub(resources, "resources")
	handle := http.FileServer(http.FS(directory))
	mux.Handle("/static/", http.StripPrefix("/static", handle))

	mux.HandleFunc("/serve-file", ServeFileEmbedHandler)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
