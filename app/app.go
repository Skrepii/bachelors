package main

import (
	"bachelors/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/login", handlers.LoginHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Listening for connections...")
}
