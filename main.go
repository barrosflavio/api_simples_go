package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Codigo que cadastra num banco de dados

    // Responde com o usuário criado!
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/user", createUser).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", router))
}
