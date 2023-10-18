package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
    Age int `json:"age"`
}

var users []User

func main() {
    r:= mux.NewRouter();

    r.HandleFunc("/users", getHandler).Methods("GET")
    r.HandleFunc("/users", postHandler).Methods("POST")
    r.HandleFunc("/users/{id}", deleteHandler).Methods("DELETE")
    r.HandleFunc("/users/{id}", putHandler).Methods("PUT")

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}

func getHandler(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(users)
}

func postHandler(w http.ResponseWriter, r *http.Request){

    var user User
    _ = json.NewDecoder(r.Body).Decode(&user)
    users = append(users, user)

    json.NewEncoder(w).Encode(users)
}


func deleteHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    intId, err := strconv.Atoi(id)

    if(err != nil){
        log.Fatal(err)
    }

    var found bool = false
    for index, user := range users {
        if user.ID == intId {
            users = append(users[:index], users[index+1:]...)
            found = true
            break
        }
    }

    if(found){
        fmt.Fprintf(w, "User deleted")
    }else{
        fmt.Fprintf(w, "User not found")
    }
}

func putHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    intId, err := strconv.Atoi(id)

    if(err != nil){
        log.Fatal(err)
    }

    var found bool = false
    for index, user := range users {
        if user.ID == intId {
            var user User
            _ = json.NewDecoder(r.Body).Decode(&user)
            users[index]= user
            found = true
            break
        }
    }

    if(found){
        json.NewEncoder(w).Encode(users)
    }else{
        fmt.Fprintf(w, "User not found")
    }

}
