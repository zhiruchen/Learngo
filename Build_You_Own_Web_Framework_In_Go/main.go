package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"encoding/json"
	"github.com/justinas/alice"
	"github.com/gorilla/context"
)

func main() {
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)

	http.Handle("/admin", commonHandlers.Append(authHandler).ThenFunc(adminHandler))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "you are on about page!")
}


// Logging middleware
func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

// recover middleware
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic: %+v\n", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// authentication middleware
func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header().Get("Authorization")
		user, err := getUser(authToken)

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		context.Set(r, "user", user)
		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}

//
func adminHandler(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	json.NewEncoder(w).Encode(user)

}