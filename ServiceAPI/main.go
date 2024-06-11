package main

import (
	"net/http"
)

func APIService() {
	http.HandleFunc("/test-connect", testConnectAPI)
	http.HandleFunc("/image/", imageHandler)
}

func main() {

	APIService()
	http.ListenAndServe(HOST_ADDRESS, nil)

}
