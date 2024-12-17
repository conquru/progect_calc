package application

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/conquru/progect_calc/pkg/calculator"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
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

func (a *Application) RunServer() error {
	http.HandleFunc("/", Calculate)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
