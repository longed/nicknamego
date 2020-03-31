package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ApiResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    ApiData `json:"data"`
}

type ApiData struct {
	Nn string `json:"nn"`
}

func startServer(ipPort string) {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/nng/v1/generate", generate)

	log.Fatal(http.ListenAndServe(ipPort, nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is nicknamego.")
}

func generate(w http.ResponseWriter, r *http.Request) {
	apiResponse := defaultApiResponse()
	switch r.Method {
	case "POST":
		if err := r.ParseMultipartForm(0); err != nil {
			fmt.Fprintf(w, "ParseForm() err:%s", err)
			return
		}
		uo := r.FormValue("userOptions")
		var userOptions UserOptions
		if err := json.Unmarshal([]byte(uo), &userOptions); err != nil {
			panic(err)
		}
		apiResponse.Data.Nn = nickname(userOptions)
		b, _ := json.Marshal(apiResponse)
		fmt.Fprintf(w, string(b))
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func defaultApiResponse() ApiResponse {
	var apiResponse ApiResponse
	apiResponse.Code = 0
	apiResponse.Message = ""

	var apiData ApiData
	apiData.Nn = ""
	return apiResponse
}
