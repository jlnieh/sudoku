package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jlnieh/sudoku/pkg/sudoku"
)

var (
	listenAddr string
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	http.HandleFunc("/solve", func(w http.ResponseWriter, r *http.Request) {
		result := struct {
			Solved bool              `json:"solved"`
			Values sudoku.ValuesType `json:"values"`
		}{false, nil}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		result.Values = sudoku.Solve(r.FormValue("grid"))
		result.Solved = (result.Values != nil)

		retObj, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(retObj)
	})

	fs := http.FileServer(http.Dir("../../web/static/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(listenAddr, nil)
}
