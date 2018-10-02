package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Status struct {
	User string  `json:"user_id"`
	Peso float32 `json:"peso"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/atualizar", trataResposta).Methods("POST")
	log.Fatal(http.ListenAndServe(":4567", router))
}

func trataResposta(w http.ResponseWriter, r *http.Request) {
	var status Status
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(string(body)), &status)

	if err != nil {
		panic(err)
	}

	fmt.Println(status.Peso)
	fmt.Println(status.User)
}
