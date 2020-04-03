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
	w.Header().Set("Content-Type", "application/json")
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
		fmt.Fprintf(w, stringApiResponse(apiResponse))
	default:
		fmt.Fprintf(w, stringApiResponse(errorApiResponse(*UnsupportRestMethod)))
	}
}

func defaultApiResponse() ApiResponse {
	var apiResponse ApiResponse
	apiResponse.Code = Success.Code
	apiResponse.Message = Success.Msg
	apiResponse.Data = defaultApiData()
	return apiResponse
}

func defaultApiData() ApiData {
	var apiData ApiData
	apiData.Nn = ""
	return apiData
}

func errorApiResponse(item Item) ApiResponse {
	var apiResponse ApiResponse
	apiResponse.Code = item.Code
	apiResponse.Message = item.Msg
	return apiResponse
}

func stringApiResponse(apiResponse ApiResponse) string {
	b, err := json.Marshal(apiResponse)
	if err != nil {
		// TODO log
		return ""
	} else {
		return string(b)
	}
}
