package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Result struct {
	Request     string
	Environment map[string]string
	Args        []string
}

func Home(w http.ResponseWriter, r *http.Request) {
	request := &bytes.Buffer{}
	_, err := io.Copy(request, r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m := map[string]string{}
	for _, s := range os.Environ() {
		splits := strings.Split(s, "=")
		m[splits[0]] = splits[1]
	}

	result := Result{
		Request:     request.String(),
		Environment: m,
		Args:        os.Args,
	}

	response, err := json.Marshal(result)
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
