package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dixonwille/wmenu/v5"
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

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(opts); return nil })

	menu.Option("List US stock marketcaps", 0, true, nil)
	menu.Option("List EU stock marketcaps", 1, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}

}

func handleFunc(opts []wmenu.Opt) {

	ListDetailsEU()

	// switch opts[0].Value {

	// case 0:
	// 	fmt.Println("Listing US stock marketcaps from Poligon.io ⌛️")
	// 	ListDetailsUS()
	// case 1:
	// 	fmt.Println("Listing EU stock marketcaps from Euronext ⌛️")
	// 	ListDetailsEU()
	// case 2:
	// 	fmt.Println("Quitting application")
	// }
}

func ListDetailsUS() {

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

	for _, ticker := range tickers {
		printDetails(ticker)
	}
}

func ListDetailsEU() {

	tickers := []string{
		"NKE",
		"TJX",
	}

	for _, ticker := range tickers {
		printDetailsEU(ticker)
	}
}

func printDetailsEU(ticker string) {
	fmt.Printf("Fetching the marketcap for %v ⌛️\n", ticker)
	details, err := api.GetDetailsEU(strings.ToUpper(ticker))

	if err != nil {
		fmt.Println("We coudn't get the details. Error details:")
		fmt.Println(err)
	} else {
		fmt.Printf("The current marketcap of %v is %.0f %v, full name: %v and website: %v \n", details.Ticker, details.MarketCap, strings.ToUpper(details.CurrencyName), details.Name, details.HomepageURL)
	}
}

func printDetails(ticker string) {
	fmt.Printf("Fetching the marketcap for %v ⌛️\n", ticker)
	details, err := api.GetDetails(strings.ToUpper(ticker))

	if err != nil {
		fmt.Println("We coudn't get the details. Error details:")
		fmt.Println(err)
	} else {
		fmt.Printf("The current marketcap of %v is %.0f %v, full name: %v and website: %v \n", details.Ticker, details.MarketCap, strings.ToUpper(details.CurrencyName), details.Name, details.HomepageURL)
	}
}
