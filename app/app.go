package main

import (
	"bachelor/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", handlers.LoginHandler)

	err := http.ListenAndServe(":3000", nil); if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Listening on port 3000...")
}