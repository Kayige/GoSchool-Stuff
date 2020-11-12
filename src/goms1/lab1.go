package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// // People struct to store json
// type People struct {
// 	Firstname string
// 	Lastname  string
// 	Details   struct {
// 		Height float32
// 		Weight float32
// 	}
// }

// Rates struct
// type Rates struct {
// 	Base   string `json:"base currency"`
// 	Symbol string `json:"destination currency"`
// }

// // Result struct
// type Result struct {
// 	Success   bool
// 	Timestamp int
// 	Base      string
// 	Date      string
// 	Rates     map[string]float64
// }

// // Error Struct
// type Error struct {
// 	Success bool
// 	Error   struct {
// 		Code int
// 		Type string
// 		Info string
// 	}
// }

var apis map[int]string

// channel c store map[int] empty interface
var c chan map[int]interface{}

func fetchData(API int) {
	url := apis[API]

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {

			var result map[string]interface{}

			json.Unmarshal(body, &result)

			var re = make(map[int]interface{})

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
					// fmt.Println(result["rates"].(map[string]interface{})["USD"])
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
					// fmt.Println(result["error"].(map[string]interface{})["info"])
				}
				c <- re

			case 2:
				if result["main"] != nil {

					kelvin := result["main"].(map[string]interface{})["temp"].(float64)
					cel := kelvin - 273.15
					re[API] = cel
					// fmt.Println(cel)
				} else {
					re[API] = result["message"]
					// fmt.Println(result["message"])
				}
				c <- re
			case 3:
				if result["status"] == "ok" {
					for _, v := range result["articles"].([]interface{}) {
						fmt.Println(v.(map[string]interface{})["source"].(map[string]interface{})["name"])
						fmt.Println(v.(map[string]interface{})["title"])
						fmt.Println(v.(map[string]interface{})["description"])
					}

				} else {
					fmt.Println("error")
				}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func lab() {
	// instantiate channel
	c = make(chan map[int]interface{})
	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=998faf77d1e6eb02cd5be046adea275d&format=1"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=feb9d4ec04b4cd5706bba29bf5c4e368"
	apis[3] = "https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=68abed74b6aa468aa1c2bcdd04ccaca9"
	go fetchData(1)
	// time.Sleep(200 * time.Millisecond)
	go fetchData(2)
	// time.Sleep(200 * time.Millisecond)
	// fmt.Scanln()
	go fetchData(3)

	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
		fmt.Printf("Result for API %v Stored!\n", i+1)
	}

	fmt.Println("Done")
}

// func main() {
//curl := "http://data.fixer.io/api/latest?access_key=998faf77d1e6eb02cd5be046adea275d&format=1"
// if resp, err := http.Get(url); err == nil {
// 	defer resp.Body.Close() // after everything is done close the connection. if sucess then close.
// 	if body, err := ioutil.ReadAll(resp.Body); err == nil {
// 		// fmt.Println(string(body))
// 		var result Result

// 		json.Unmarshal(body, &result)
// 		if result.Success { // managed to fetch result
// 			keys := make([]string, 0, len(result.Rates))
// 			for k := range result.Rates {
// 				keys = append(keys, k)
// 			}
// 			sort.Strings(keys)
// 			for _, k := range keys {
// 				fmt.Println(k, result.Rates[k])
// 			}
// 		} else { // error
// 			var error Error
// 			json.Unmarshal(body, &error)
// 			fmt.Println("Code: ", error.Error.Code)
// 			fmt.Println("Info: ", error.Error.Info)
// 			fmt.Println("Type: ", error.Error.Type)

// 		}
// 	} else {
// 		log.Fatal(err)
// 	}

// } else {
// 	log.Fatal(err)
// }
// }

// func main() {
// jsonString := `[{
// "firstname":"John", "lastname":"Doe",
// "details": {
// 		"height":175, "weight":70.0
// }
// },
// {
// "firstname":"Mary", "lastname":"Doe",
// "details": {
// 		"height":105, "weight":71.1
// }
// }]`

// jsonString2 := `{
//  "base currency": "EUR",
//  "destination currency": "USD"
//  }`

// var person []People
// err := json.Unmarshal([]byte(jsonString), &person)
// fmt.Println(person)
// for _, el := range person {
// 	fmt.Println(el.Firstname, el.Lastname)
// 	fmt.Println(el.Details.Height, el.Details.Weight)
// }
// fmt.Println("jsonString:", jsonString)
// fmt.Println("Error:", err)

// jsonString3 :=
// 	`{
//     "success": true,
//     "timestamp": 1588779306,
//     "base": "EUR",
//     "date": "2020-05-06",
//     "rates": {
//         "AUD": 1.683349,
//         "CAD": 1.528643,
//         "GBP": 0.874757,
//         "SGD": 1.534513,
//         "USD": 1.080054
//     }
// }`

// empty interface the value can be anything
// var result map[string]interface{}
// json.Unmarshal([]byte(jsonString3), &result)
// fmt.Println(result["success"])
// currRates := result["rates"] // value is an interface{}
// fmt.Println(currRates)

// initialize a variable and create another interface{}
// SGD := currRates.(map[string]interface{})["SGD"]
// fmt.Println(SGD)

// var rates Rates
// json.Unmarshal([]byte(jsonString2), &rates)
// // fmt.Println("jsonString2:", jsonString2)
// fmt.Println(rates.Base)
// fmt.Println(rates.Symbol)
// }
