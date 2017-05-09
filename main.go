package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"google.golang.org/appengine"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Slow endpoint")

	r := mux.NewRouter()
	r.HandleFunc("/{seconds}", imSlow).Methods("GET")

	// Non-appengine
	//http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	//http.ListenAndServe(":12001", nil)

	// AppEngine
	// AppEngine / Compute Engine Health check endpoint
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	appengine.Main()
}

func imSlow(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var seconds int64

	secstr := mux.Vars(r)["seconds"]
	seconds, err := strconv.ParseInt(secstr, 10, 64)
	if err != nil {
		seconds = 0
	}
	if seconds > 20 {
		seconds = 20
	}
	log.Printf("Requesting a delay of %s seconds, using %v", secstr, seconds)

	time.Sleep(time.Duration(seconds) * time.Second)

	elapsed := time.Since(start)

	response := struct {
		Elapsed   string `json:"elapsed"`
		Requested string `json:"requested"`
	}{
		Elapsed:   elapsed.String(),
		Requested: secstr,
	}

	d, err := json.Marshal(&response)
	if err != nil {
		http.Error(w, "can't even", http.StatusInternalServerError)
	}
	w.Write(d)
}
