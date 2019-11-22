package service

import (
	"net/http"
	"os"
	"text/template"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/unknown", NotImplemented).Methods("GET")
	mx.HandleFunc("/", home).Methods("GET")
	mx.HandleFunc("/", login).Methods("POST")
	
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
}

func login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	t:=template.Must(template.ParseFiles("templates/info.html"))
	t.Execute(w,map[string] string{
		"username": username,
		"password": password,
	})
}

func home(w http.ResponseWriter, r *http.Request)  {

	t := template.Must(template.ParseFiles("templates/login.html"))
	t.Execute(w, nil)
}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			WELCOME      string `json:"WELCOME"`
			
		}{WELCOME: "Hello"})
}
}

// NotImplemented 
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", 501)
}