package application

import (
	"encoding/json"
	"net/http"

	"github.com/conquru/project_calc/pkg/calculator"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

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
		response := Response{Result: 0, Error: "Expression is notvalid"}
		http.StatusText(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := calculator.Calc(request.Expression)
	if err != nil {
		response := Response{Result: 0, Error: "Expression is notvalid"}
		http.StatusText(422)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{Result: result, Error: "invalid"}
	http.StatusText(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", Calculate)
	return http.ListenAndServe(":8080", nil)
}
