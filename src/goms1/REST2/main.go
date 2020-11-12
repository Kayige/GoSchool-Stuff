package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:5001/api/v1/courses"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

func getCourse(code string) {
	url := baseURL
	if code != "" {
		url = baseURL + "/" + code + "?key=" + key

	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func addCourse(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(baseURL+"/"+code+"?key="+key, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Println("Error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
	}
}

func updateCourse(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest(http.MethodPut, baseURL+"/"+code+"?key="+key, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()

	}

}

func deleteCourse(code string) {
	request, err := http.NewRequest(http.MethodDelete, baseURL+"/"+code+"?key="+key, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()

	}
}
func main() {
	// jsonData := map[string]string{"title": "Applied Go Programming"}
	jsonData := map[string]string{"title": "Android Programming"}
	addCourse("IOS101", jsonData)
	getCourse("")

}
