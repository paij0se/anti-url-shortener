package routes

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bruh-boys/anti-url-shortener/src/controllers"
	"github.com/gorilla/mux"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-limiter/memorystore"
)

var (
	store, _ = memorystore.New(&memorystore.Config{
		// Number of tokens allowed per interval.
		Tokens: 15,

		// Interval until tokens reset.
		Interval: time.Minute,
	})
)

func SetupRoutes() error {

	mux.NewRouter()
	r := mux.NewRouter()

	middleware, _ := httplimit.NewMiddleware(store, httplimit.IPKeyFunc())
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("src/view/index.html")
		t.Execute(rw, map[string]string{"url": ""})

	}).Methods("GET")
	r.HandleFunc(`/static/{file:[\w\W\/]+}`, func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path[1:])
		http.ServeFile(w, r, "src"+r.URL.Path)
	}).Methods("GET")
	r.Handle("/no-bitly", middleware.Handle(http.HandlerFunc(controllers.NoBitly))).Methods("POST")

	port, err := os.LookupEnv("PORT")
	if !err {
		port = "8080"
	}
	log.Println("Listening on port: " + port)
	return http.ListenAndServe(":"+port, r)

}
