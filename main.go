package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fuww/top200/api"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	tickers := []string{
		"NKE",
		"TJX",
		"VFC",
		// "COH",
		"GPS",
		// "LB",
		// "TIF",
		"PVH",
		// "KORS",
		"M",
		"RL",
		"JWN",
		"HBI",
		"UA",
		"CRI",
		"FL",
		"COLM",
		"LULU",
		"URBN",
		"AEO",
		"DKS",
	}

	// Iterate over the slice using a for loop
	for _, ticker := range tickers {
		// details, err := api.GetDetails(ticker)
		// print(details, err)
		printDetails(ticker)
	}

}

func printDetails(ticker string) {
	fmt.Printf("Fetching the marketcap for %v ⌛️\n", ticker)
	details, err := api.GetDetails(strings.ToUpper(ticker))

	if err != nil {
		fmt.Println("We coudn't get the details. Details:")
		fmt.Println(err)
	} else {
		fmt.Printf("The current marketcap of %v is %.0f %v, full name: %v and website: %v LastUpdatedUTC: %v \n", details.Ticker, details.MarketCap, strings.ToUpper(details.CurrencyName), details.Name, details.HomepageURL, details.LastUpdatedUTC)
	}
}
