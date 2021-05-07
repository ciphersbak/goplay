package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func main() {
	url := "https://ipinfo.io/?token=YOUR_TOKEN"
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
	var data map[string]interface{}
	// data := map[string]interface{}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	for k, v := range data {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}
	value, ok := data["ip"]
	if ok {
		fmt.Printf("Yes, IP has a value and it is %q\n", value)
	}
	// list all kvp
	for key, value := range data {
		fmt.Println(key, value)
	}
	// list only key
	for key := range data {
		fmt.Println(key)
	}
	// list only value
	for _, value := range data {
		fmt.Println(value)
	}
	// store keys and iterate alphabetically
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Println("In alphabetical order")
	for _, value := range keys {
		fmt.Println(value, data[value])
	}
}
