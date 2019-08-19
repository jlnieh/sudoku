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
	listenAddr  string
	webpagePath string
)

func httpsrv() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("HTTP Server(%s) is starting at (%s)...\n", webpagePath, listenAddr)

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

	fs := http.FileServer(http.Dir(webpagePath))
	http.Handle("/", http.StripPrefix("/", fs))

	logger.Fatal(http.ListenAndServe(listenAddr, nil))
}

func main() {
	flag.StringVar(&webpagePath, "webpagepath", "web/static", "Static web pages path.")
	flag.StringVar(&listenAddr, "listenaddr", ":8080", "server listen address")
	flag.Parse()

	if flag.NArg() > 0 {
		result := sudoku.Solve(flag.Arg(0))
		if result == nil {
			fmt.Println("Sorry! Cannot solve the puzzle!")
		} else {
			sudoku.Display(result)
		}
		return
	}

	httpsrv()
}
