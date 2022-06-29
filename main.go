package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"ws/entity"
	"ws/service"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

var users = map[int]entity.User{
	1: {
		Id:        1,
		Username:  "Delon",
		Email:     "email",
		Password:  "password",
		Age:       13,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	2: {
		Id:        2,
		Username:  "Chandra",
		Email:     "email",
		Password:  "password",
		Age:       123,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	3: {
		Id:        3,
		Username:  "Bambang",
		Email:     "email",
		Password:  "password",
		Age:       18,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/register", register)
	r.HandleFunc("/user", UserHandler)
	r.HandleFunc("/user/{Id}", UserHandler)
	r.HandleFunc("/user/{Id}", UserHandler)

	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world"
	fmt.Fprint(w, msg)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var user entity.User
		if err := decoder.Decode(&user); err != nil {
			w.Write([]byte("error decoding json body"))
			return
		}

		userSvc := service.NewUserService()
		userTemp := userSvc.Register(&user)

		jsonData, _ := json.Marshal(userTemp)

		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id := params["Id"]
	if r.Method == "GET" {
		if Id != "" {
			tempId, _ := strconv.Atoi(Id)
			GetUserById(w, r, tempId)
		} else {
			GetAllUser(w, r)
		}
	}
	if r.Method == "DELETE" {
		tempId, _ := strconv.Atoi(Id)
		DeleteUser(w, r, tempId)
	}
	if r.Method == "POST" {
		AddUser(w, r)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	test, _ := json.Marshal(users[id])
	w.Header().Add("Content-Type", "application/json")
	w.Write(test)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	values := []entity.User{}
	for _, value := range users {
		values = append(values, value)
	}
	test, _ := json.Marshal(values)
	w.Header().Add("Content-Type", "application/json")
	w.Write(test)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	for _, value := range users {
		if value.Id == id {
			delete(users, value.Id)
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("User deleted successfully"))
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	users[len(users)+1] = user
}
