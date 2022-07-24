package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	var directory string
	port := "8000"

	if len(args) == 0 {
		directory, _ = filepath.Abs(".")
	} else if len(args) == 1 {
		directory, _ = filepath.Abs(args[0])
	} else {
		directory, _ = filepath.Abs(args[0])
		port = args[1]
	}

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving \u001b[32m%s\u001b[0m on HTTP port: \u001b[34m%s\u001b[0m\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
