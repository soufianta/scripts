package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("C:\\Users\\Sousou86\\SCRIPTS")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Fprintln(w, file.Type(), file.Name())
	}
}
func main() {
	http.HandleFunc("/files", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
