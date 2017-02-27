package main

import (
	"fmt"
	"net/http"
)

//this will always start the server on localhost and make it
//accessible from outside the container
func startServer(port int) {
	fmt.Println("starting url2pdf microserver...")
	http.HandleFunc("/", UrlToPdfService)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
