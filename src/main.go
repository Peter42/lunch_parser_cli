package main

import "fmt"
import "net/http"
import "encoding/json"
import "os"

type Response struct {
	Menus          []LunchMenu   `json:"menus"`
	GenerationTime int           `json:"generationTime"`
	MenuForDay     int           `json:"menuForDay"`
}
type LunchMenu struct {
	Name       string            `json:"name"`
	LunchItems []LunchMenuItem   `json:"lunchItems"`
}

type LunchMenuItem struct {
	ItemName   string            `json:"itemName"`
}


func main() {
	resp, err := http.Get("https://temp.philipp1994.de/sap/lunch/api/v1/")
	
	if err != nil {
		fmt.Print("Could not connect to API: ")
		fmt.Println(err)
		os.Exit(3)
	}
	
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	var res Response
	err = dec.Decode(&res)
	if err != nil {
		fmt.Print("Could not parse json: ")
		fmt.Println(err)
		os.Exit(3)
	}
	
	for _, v := range res.Menus {
		fmt.Println(v.Name)
		for _, v := range v.LunchItems {
			fmt.Print("\t")
			fmt.Println(v.ItemName)
		}
	}
}