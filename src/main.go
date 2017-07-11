package main

import "fmt"
import "net/http"
import "encoding/json"
import "os"
import "unicode/utf8"

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
	Comment    string            `json:"comment"`
	Price      float64           `json:"price"`
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
	
	var MaxItemNameLength int
	MaxItemNameLength = 0
	var MaxCommentLength int
	MaxCommentLength = 0
	
	for _, v := range res.Menus {
		for _, v := range v.LunchItems {
			
			ItemNameLength := utf8.RuneCountInString(v.ItemName)
			CommentLength  := utf8.RuneCountInString(v.Comment)
			
			if(ItemNameLength > MaxItemNameLength) {
				MaxItemNameLength = ItemNameLength
			}
			
			if(CommentLength > MaxCommentLength) {
				MaxCommentLength = CommentLength
			}
		}
	}
	
	for _, v := range res.Menus {
		fmt.Println(v.Name)
		for _, v := range v.LunchItems {
			fmt.Print("\t")
			fmt.Print(v.ItemName)
			for i := utf8.RuneCountInString(v.ItemName); i <= MaxItemNameLength + 1; i++ {
				fmt.Print(" ")
			}
			
			fmt.Print(v.Comment)
			for i := utf8.RuneCountInString(v.Comment); i <= MaxCommentLength + 1; i++ {
				fmt.Print(" ")
			}
			
			if v.Price != -1 {
				fmt.Printf("%0.2f\n", v.Price)
			} else {
				fmt.Println("   u")
			}
		}
	}
}