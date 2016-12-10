package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/yehohanan7/rafty/calc"
)

type Value struct {
	X int `json:"x"`
}

func Peers() []string {
	var peers []string
	file, _ := ioutil.ReadFile("peers.json")
	json.Unmarshal(file, &peers)
	return peers
}

func Run(w http.ResponseWriter, r *http.Request, fn func(int)) {
	defer r.Body.Close()
	var value Value
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fn(value.X)
	w.Write([]byte("value updated successfuly"))
}

func main() {
	peers := Peers()
	id := os.Args[1]
	id_int, _ := strconv.Atoi(id)
	calculator := calc.NewCalculator("peer"+id, peers[id_int-1], id)
	r := mux.NewRouter()

	r.HandleFunc("/value", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("current value: %v", calculator.Value)))
	})

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		Run(w, r, calculator.Add)
	}).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":800"+id, r))
}
