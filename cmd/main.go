package main

import (
	"encoding/json"
	"log"
	"net/http"

	calculator "github.com/conquru/progect_calc/pkg/calculator"
)

// запрос
type Request struct {
	Expression string `json:"expression"`
}

// ответ
type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request Request
	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calculator.Calc(request.Expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{Result: result, Error: "nil"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/calculate", Calculate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
