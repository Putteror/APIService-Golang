package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type returnJSONData struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Desc string `json:"desc"`
}

func testConnectAPI(resp http.ResponseWriter, request *http.Request) {

	var desc_content string = ""

	fmt.Printf("Received %s request from IP: %s\n", string(request.Method), request.RemoteAddr)

	resp.Header().Set("Content-Type", "application/json")

	switch method := request.Method; method {
	case http.MethodPost:
		desc_content = "POST success"
	case http.MethodGet:
		desc_content = "GET success"
	}

	returnJSONData := returnJSONData{
		Code: 200,
		Data: "success",
		Desc: desc_content,
	}

	returnData, err := json.Marshal(returnJSONData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(resp, "Error marshalling response to JSON: %v", err)
		return
	}

	resp.Write(returnData)
}
