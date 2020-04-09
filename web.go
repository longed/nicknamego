package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"sync/atomic"
)

var (
	counter int32 = 0 // global counter
)
type ApiResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    ApiData `json:"data"`
}

type ApiData struct {
	Nn   string `json:"nn"`
	Date string `json:"date"`
}

func startServer(ipPort string) {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/nng/v1/generate", generateHandler)
	log.Fatal(http.ListenAndServe(ipPort, nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	apiResponse := defaultApiResponse()
	apiResponse.Data.Nn = "hello, this is nicknamego."
	fmt.Fprintf(w, apiResponse.toStr())
	fmt.Printf("counter: %d\n", get())
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	add()
	apiResponse := defaultApiResponse()
	switch r.Method {
	case "POST":
		if err := r.ParseMultipartForm(0); err != nil {
			fmt.Fprintf(w, errorApiResponseWithErr(ParseMultipartFormErr, err).toStr())
			return
		}
		uo := r.FormValue("userOptions")
		var userOptions UserOptions
		if err := json.Unmarshal([]byte(uo), &userOptions); err != nil {
			fmt.Fprintf(w, errorApiResponseWithErr(UnmarshalJsonErr, err).toStr())
			return
		}
		apiData := nickname(userOptions)
		apiResponse.Data = apiData
		fmt.Fprintf(w, apiResponse.toStr())
	default:
		fmt.Fprintf(w, errorApiResponse(UnsupportRestMethod).toStr())
	}
}

func add() {
	atomic.AddInt32(&counter, 1)
}

func get() int32 {
	return atomic.LoadInt32(&counter)
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
	apiData.Date = time.Now().Format("2006-01-02 15:04:05")
	return apiData
}

func errorApiResponse(item *ErrItem) ApiResponse {
	var apiResponse ApiResponse
	apiResponse.Code = item.Code
	apiResponse.Message = item.Msg
	return apiResponse
}

func errorApiResponseWithErr(item *ErrItem, err error) ApiResponse {
	apiResponse := errorApiResponse(item)
	apiResponse.Message += err.Error()
	return apiResponse
}

func (apiResponse ApiResponse) toStr() string {
	b, err := json.Marshal(apiResponse)
	if err != nil {
		// TODO log
		return ""
	} else {
		return string(b)
	}
}
