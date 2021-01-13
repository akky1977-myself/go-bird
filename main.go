package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akky/bird/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Bird Encyclopedia")

	r := newRouter()

	http.ListenAndServe(getPort(), r)
}

func getPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("cannot find .env file: ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		return ":8080"
	}

	return ":" + port
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticFileDirectory := http.Dir("./assets/")
	// /assets_path/ is uri for routing
	staticFileHandler := http.StripPrefix("/assets_path/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets_path/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/okay", Okay).Methods("GET")
	r.HandleFunc("/bird", handler.CreateBirdHandler).Methods("POST")
	return r
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Brird Encyclopedia is running.")
}

func Okay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "okay")
}
