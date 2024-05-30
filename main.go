package main

import (
	"log"
	"time"
)

var current_date = time.Now().Format("020106")
var yesterday_date = time.Now().AddDate(0, 0, -1).Format("020106")
var current_time = time.Now().Format("15:04:05")
var file_Link_format = ("https://www.bseindia.com/download/BhavCopy/Equity/EQ" + yesterday_date + "_CSV.ZIP")

func time_check() {
	var current_time_check = current_time
	if current_time_check == "" {
		log.Fatal("Error in time.")

	} else {
		if current_time != "00.00.00" {

		}

	}

}

func main() {
	time_check()

}
