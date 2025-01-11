package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Skb08/urlshort/url-short"
)

//Guidance for run
// 1.for default maphandler ./urlshort
// 2.for yaml ./urlshort -yaml urls.yaml
// 3.for json ./urlshort -json urls.json

func main() {
	yamlFile := flag.String("yaml", "", "Path to the YAML file containing URL mappings")
	jsonFile := flag.String("json", "", "Path to the JSON file containing URL mappings")
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort": "https://godoc.org/github.com/gophercises/urlshort",
		"/urlshort2":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Initialize the handler
	var handler http.Handler
	if *yamlFile != "" {
		yamlData, err := ioutil.ReadFile(*yamlFile)
		if err != nil {
			log.Fatalf("Failed to read YAML file: %v", err)
		}
		handler, err = urlshort.YAMLHandler(yamlData, mapHandler)
		if err != nil {
			log.Fatalf("Failed to create YAMLHandler: %v", err)
		}
	} else if *jsonFile != "" {
		jsonData, err := ioutil.ReadFile(*jsonFile)
		if err != nil {
			log.Fatalf("Failed to read JSON file: %v", err)
		}
		handler, err = urlshort.JSONHandler(jsonData, mapHandler)
		if err != nil {
			log.Fatalf("Failed to create JSONHandler: %v", err)
		}
	} else {
		log.Println("No YAML or JSON file provided. Using default MapHandler.")
		handler = mapHandler
	}

	// Start the server
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
