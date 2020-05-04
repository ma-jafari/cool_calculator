package main

import (
	"encoding/json"
	"net/http"
)

type numStruct struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type responseData struct {
	Add float64 `json:"add"`
	Sub float64 `json:"sub"`
	Mul float64 `json:"mul"`
	Div float64 `json:"div"`
}

func process(numsData numStruct) responseData {
	var response responseData
	response.Add = numsData.Num1 + numsData.Num2
	response.Sub = numsData.Num1 - numsData.Num2
	response.Mul = numsData.Num1 * numsData.Num2
	response.Div = numsData.Num1 / numsData.Num2

	return response
}

func cal(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var numbers numStruct
	var operations responseData

	decoder.Decode(&numbers)

	operations = process(numbers)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(operations); err != nil {
		println(err.Error())
		return
	}
}

func main() {
	println("Hello I'm a cool calculator! :D")
}
