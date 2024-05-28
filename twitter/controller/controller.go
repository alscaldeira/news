package controller

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/alscaldeira/twitter/model"
	service "github.com/alscaldeira/twitter/service"
	utils "github.com/alscaldeira/twitter/utils"
	"github.com/gorilla/mux"
)

func InitializeServer() {
	router := generateRoutes()
	utils.Log("Server initialized!")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func generateRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/twitter/user", postUser).Methods("POST")
	router.HandleFunc("/twitter/login/{username}", login).Methods("GET")

	router.HandleFunc("/twitter/post", post).Methods("POST")

	return router
}

func post(w http.ResponseWriter, r *http.Request) {
	utils.Log("Controller posting on Twitter")

	var post model.Post
	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		log.Fatal(err)
	}

	service.Post(post.Username, post.Password, post.Content)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	utils.Log("Controller posting user")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	service.PostUser(user.Username, user.Password)
}

func login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	service.Login(username)
}
