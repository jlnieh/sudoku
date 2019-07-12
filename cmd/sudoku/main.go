package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")

		values := sudoku.Solve(r.FormValue("grid"))
		if values == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Sorry! I cannot solve the puzzle!")
			return
		}

		retObj, _ := json.Marshal(values)
		w.WriteHeader(http.StatusOK)
		w.Write(retObj)
	})

	fs := http.FileServer(http.Dir("../../web/static/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(listenAddr, nil)
}
