package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
 * Path of configuration files
 */
const ConfigPath = "config/"

/**
 * Configuration file extension
 */
const FileExtension = ".file"

/**
 * Main application
 */
func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/config", list).Methods("GET")
	router.HandleFunc("/config/{id}", create).Methods("POST")
	router.HandleFunc("/config/{id}", read).Methods("GET")

	fmt.Println("Running config server at port 8080!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/**
 * Index application response
 */
func index(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintln(w, "Hello, welcome to remote config!")
}

/**
 * Create configuration
 */
func create(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	if s == "" {
		http.Error(w, "Invalid config name", 500)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s%s%s", ConfigPath, mux.Vars(r)["id"], FileExtension), b, 0644)
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}

	fmt.Fprintln(w, string(b))
}

/**
 * Read configuration
 */
func read(w http.ResponseWriter, r *http.Request)  {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s%s%s", ConfigPath, mux.Vars(r)["id"], FileExtension))
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}

	fmt.Fprintln(w, string(b))
}

/**
 * List all configurations
 */
func list(w http.ResponseWriter, r *http.Request){
	files, err := ioutil.ReadDir(ConfigPath)
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}

	for _, f := range files {
		fmt.Fprintln(w, strings.Replace(f.Name(), FileExtension, "", -1))
	}
}
