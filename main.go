package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jlnieh/sudoku/pkg/sudoku"
	"github.com/spf13/viper"
)

var (
	configFile  string
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

	http.ListenAndServe(listenAddr, nil)
}

func retrieveConfigs() {
	viper.SetDefault("ListenAddr", ":8080")
	viper.SetDefault("WebpagePath", "web/static")

	if configFile != "" {
		content, err := ioutil.ReadFile(configFile)

		if err != nil {
			panic(fmt.Errorf("Fatal error to read config file(%s): %s", configFile, err))
		}
		viper.ReadConfig(bytes.NewBuffer(content))
	} else {
		viper.SetConfigName("config")        // name of config file (without extension)
		viper.AddConfigPath("/etc/sudoku/")  // path to look for the config file in
		viper.AddConfigPath("$HOME/.sudoku") // call multiple times to add many search paths
		viper.AddConfigPath(".")             // optionally look for config in the working directory

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				// fmt.Printf("%s\n", err)
			} else {
				panic(fmt.Errorf("Fatal error to parse config file: %s", err))
			}
		}

	}

	if webpagePath == "" {
		webpagePath = viper.GetString("WebpagePath")
	}

	if listenAddr == "" {
		listenAddr = viper.GetString("ListenAddr")
	}
}

func main() {
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&webpagePath, "webpagepath", "", "Static web pages path.")
	flag.StringVar(&listenAddr, "listenaddr", "", "server listen address")
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

	if (webpagePath == "") && (listenAddr == "") {
		retrieveConfigs()
	}

	httpsrv()
}
