package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/callback", callback)
	fmt.Println("Callback listening on http://<device_ip>:4444/callback POST")
	http.ListenAndServe("0.0.0.0:4444", nil)
}
