package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	HOST = ""
	PORT = 8080
)

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", HOST, PORT),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerValues := make(map[string][]string)
			for key, value := range r.Header {
				headerValues[key] = value
			}

			queryValues := make(map[string][]string)
			for key, value := range r.URL.Query() {
				queryValues[key] = value
			}

			requestAddress := r.RemoteAddr

			method := r.Method

			host := r.Host

			body := r.Body

			requestData := map[string]interface{}{
				"host":         host,
				"method":       method,
				"address":      requestAddress,
				"headers":      headerValues,
				"body":         body,
				"query_params": queryValues,
			}

			switch r.Method {
			case "POST":
				r.ParseForm()
				formValues := make(map[string][]string)
				for key, value := range r.PostForm {
					formValues[key] = value
				}

				requestData["form_params"] = formValues

			case "PUT":
				r.ParseForm()
				formValues := make(map[string][]string)
				for key, value := range r.PostForm {
					formValues[key] = value
				}

				requestData["form_params"] = formValues

			case "PATCH":
				r.ParseForm()
				formValues := make(map[string][]string)
				for key, value := range r.PostForm {
					formValues[key] = value
				}

				requestData["form_params"] = formValues
			}

			jsonString, err := json.Marshal(requestData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fmt.Println(string(jsonString))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonString)
		}),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed to start")
	}
}
