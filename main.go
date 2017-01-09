package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("error ListenAndServe:", err)
	}
}

func root(w http.ResponseWriter, req *http.Request) {
	contents, err := ioutil.ReadFile("/usr/local/tmp.txt")
	if err != nil {
		fmt.Fprintln(w, "error reading file:", err)
		return
	}
	fmt.Fprintln(w, "Content:", string(contents))
}
