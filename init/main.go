package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server", err)
		panic(err)
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
