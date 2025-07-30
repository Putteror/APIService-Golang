package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type returnJSONData struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func callback(resp http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(resp, "Method %s not allowed", request.Method)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(resp, "Error reading request body: %v", err)
		return
	}

	currentTime := time.Now()
	fmt.Println("Default timestamp:", currentTime.Unix())
	fmt.Println("Formatted timestamp:", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Received POST request body:\n%s\n", string(body))

	resp.Header().Set("Content-Type", "application/json")

	returnJSONData := returnJSONData{
		Code: 200,
		Data: "success",
	}

	returnData, err := json.Marshal(returnJSONData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Error marshalling response to JSON: %v", err)
		return
	}

	resp.Write(returnData)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the port number: ")
	portStr, _ := reader.ReadString('\n')
	portStr = strings.TrimSpace(portStr) // Remove leading/trailing whitespace, including newline

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port number. Please enter a number.")
		return
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)

	http.HandleFunc("/callback", callback)
	fmt.Printf("Callback listening on http://<device_ip>:%d/callback POST\n", port)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
