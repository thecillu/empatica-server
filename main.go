package main

import (
	"empatica-server/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/articles", saveArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles", manageOptions).Methods("OPTIONS")
	myRouter.HandleFunc("/articles/{id}", manageOptions).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":80", myRouter))
}

func manageOptions(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.WriteHeader(http.StatusOK)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage v.0.0.4!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var articles []model.Article
	var err error
	articles, err = model.GetAllArticles()
	if err != nil {
		http.Error(w, "Ops...sorry but we have some problems", http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(articles)
	}
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	var article model.Article
	var err error
	article, err = model.GetArticle(id)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(article)
	}
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article model.Article
	json.Unmarshal(reqBody, &article)
	var err error
	err = model.SaveArticle(article)
	if err != nil {
		http.Error(w, "Ops...sorry but we have some problems", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	id := vars["id"]
	var err error
	err = model.DeleteArticle(id)
	if err != nil {
		http.Error(w, "Ops...sorry but we have some problems", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article model.Article
	json.Unmarshal(reqBody, &article)
	var err error
	var result int
	result, err = model.UpdateArticle(article)
	if err != nil {
		http.Error(w, "Ops...sorry but we have some problems", http.StatusInternalServerError)
	} else {
		w.WriteHeader(result)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin,  Pragma, Cache-Control, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,  X-Auth-Token, X-Requested-With, Access-Control-Allow-Headers")
}

func main() {
	var DB_ENDPOINT string
	DB_ENDPOINT = os.Getenv("DB_ENDPOINT")
	var DB_USERNAME string
	DB_USERNAME = os.Getenv("DB_USERNAME")
	var DB_PASSWORD string
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME string
	DB_NAME = os.Getenv("DB_NAME")
	var DB_CONNECTION_STRING string
	DB_CONNECTION_STRING = DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_ENDPOINT + ")/" + DB_NAME
	fmt.Println(DB_CONNECTION_STRING)
	m, err := migrate.New(
		"file://db/migrations",
		"mysql://"+DB_CONNECTION_STRING)
	m.Steps(1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Empatica Rest API")
	handleRequests()
}
