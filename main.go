package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cohix/generator"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable must be set")
	}

	root := os.Getenv("ROOT_DIR")
	if root == "" {
		log.Fatal("ROOT_DIR environment variable must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving /")
		fmt.Fprintf(w, generator.GenerateIndexPageFromRootDir(root))
	})

	log.Println("MedServ listening on :" + port + "\n\n")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
