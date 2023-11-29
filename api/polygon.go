package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func GetDetails(ticker string) (*Details, error) {
	if len(ticker) == 0 {
		return nil, errors.New("ticker empty")
	}

	POLYGON_API_KEY := os.Getenv("POLYGON_API_KEY")
	date := time.Date(2023, 11, 01, 0, 0, 0, 0, time.Local)

	// init client
	c := polygon.New(POLYGON_API_KEY)

	params := models.GetTickerDetailsParams{
		Ticker: ticker,
	}.WithDate(models.Date(date))

	// make request
	r, err := c.GetTickerDetails(context.Background(), params)
	details := Details{}
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(r.Results.MarketCap)
	fmt.Println(r.Results.LastUpdatedUTC)

	jsonBlob, err := json.Marshal(r.Results)
	if err != nil {
		return nil, err
	}

	var result Results // my own
	// var result = models.GetTickerDetailsResponse.Results // from polygon

	err = json.Unmarshal(jsonBlob, &result)
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Printf("%+v", result)

	details.Ticker = ticker
	details.MarketCap = result.MarketCap
	details.Name = result.Name
	details.CurrencyName = result.CurrencyName
	details.CurrencySymbol = result.CurrencySymbol
	details.Active = result.Active
	// details.BaseCurrencyName = result.BaseCurrencyName
	// details.BaseCurrencySymbol = result.BaseCurrencySymbol
	details.Description = result.Description
	details.HomepageURL = result.HomepageURL
	details.WeightedSharesOutstanding = result.WeightedSharesOutstanding
	// marketcap := result.MarketCap
	// details.MarketCap = marketcap
	// name := result.Name
	// details.Name = name

	// fmt.Printf("ticker: %v name: %v marketcap: %.f \n", details.Ticker, details.Name, details.MarketCap)

	if err != nil {
		return nil, err
	}

	// TODO: create store to file function
	// log.Print(r) // do something with the result
	// // log.Print(r.Results.Ticker)

	// filename := r.Results.Ticker + "_" + date.Format("2006-01-02") + "_" + ".json"

	// f, e := os.Create("output/" + filename)
	// if e != nil {
	// 	panic(e)
	// }
	// defer f.Close()

	// jsonByte, err := json.Marshal(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// reader := bytes.NewReader(jsonByte)
	// io.Copy(f, reader)

	// details = string(jsonByte)

	return &details, nil

}
