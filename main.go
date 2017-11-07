package main

import (
	"net/http"
	"time"
	"encoding/json"
	"database/sql"
	"uk_bankholidays/holidays"
	"uk_bankholidays/database"
	"log"
)

const URL ="https://www.gov.uk/bank-holidays.json"

func main(){


	bank_holidays := holidays.Country{}

	// Download & parse data from URL
	var cli = &http.Client{Timeout: 10 * time.Second}

	r, err := cli.Get(URL)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&bank_holidays)
	if err != nil {
		panic(err)
	}
	// -------------------------------

	// Connect to database
	db,err := database.DatabaseConnect()
	defer db.Close()

	if err!=nil {
		log.Fatal(err)
	}
	// -------------------------------


	// loop through and save country dates
	if len(bank_holidays.England.Events)>0{
		for _,holiday := range bank_holidays.England.Events {
			SaveHoliday(holiday,"England",db)
		}
	}

	if len(bank_holidays.Scotland.Events)>0{
		for _,holiday := range bank_holidays.England.Events {
			SaveHoliday(holiday,"Scotland",db)
		}
	}

	if len(bank_holidays.NIreland.Events)>0{
		for _,holiday := range bank_holidays.England.Events {
			SaveHoliday(holiday,"N Ireland",db)
		}
	}
	// -------------------------------

}

func SaveHoliday(h holidays.Holiday, country string, db *sql.DB) (err error){

	_, err = db.Exec("INSERT INTO bank_holidays (description,event_date,country,entered) VALUES (?,?,?,NOW())",h.Title,h.Date,country)
	return err

}


