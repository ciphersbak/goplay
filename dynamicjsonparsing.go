package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://ipinfo.io/?token=9acfb367e2154c"
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

}
