package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/zabawaba99/firego.v1"

	"github.com/gorilla/mux"
)

type Status struct {
	User        string `json:"user_id"`
	Peso        string `json:"peso"`
	Porcentagem int    `json:"porcentagem"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/atualizar", trataResposta).Methods("POST")
	log.Fatal(http.ListenAndServe(":4567", router))
}

func trataResposta(w http.ResponseWriter, r *http.Request) {
	fireInstance := firego.New("https://gasnow-acf5c.firebaseio.com/", nil)
	var status Status
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal([]byte(string(body)), &status)
	if err != nil {
		panic(err)
	}
	status.Porcentagem = calcularPorcentagem(status.Peso)
	db, _ := fireInstance.Ref("/botijoes")

	var result map[string]interface{} = make(map[string]interface{})

	err = db.OrderBy("user_id").EqualTo("1").LimitToFirst(1).Value(&result)

	for k, _ := range result {
		db, _ = db.Ref("/botijoes/" + k)
		status.Porcentagem = 100000
		db.Update(status)
	}

}

func calcularPorcentagem(peso string) int {
	//CONVERTER O PESO PARA DOUBLE
	//CALCULAR A PORCENTAGEM
	return 0
}
