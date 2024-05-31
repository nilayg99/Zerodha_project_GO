package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

var current_date = time.Now().Format("020106")
var yesterday_date = time.Now().AddDate(0, 0, -1).Format("020106")
var current_time = time.Now().Format("15:04:05")
var file_Link_format = ("https://www.bseindia.com/download/BhavCopy/Equity/EQ" + yesterday_date + "_CSV.ZIP")
var file_name = ("EQ" + yesterday_date + "_CSV.ZIP")
var BSE_file_path = "D:/Project/Project_New/Zerodha_project/Zerodha_project_GO/BSE_File_Saved"

/*
	func time_check() {
		var current_time_check = current_time
		if current_time_check == "" {
			log.Fatal("Error in time.")

		} else {
			if current_time != "00.00.00" {
				var current_time string
				var timeLeft = time.Until(current_time)

			}

		}

}
*/
// file_exist checks if the file exists at the given path.
func file_exist(file_name string) bool {
	_, err := os.Stat(file_name)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
func URL_status_check() {
	resp, err := http.Get(file_Link_format)
	if err != nil {
		log.Fatal(err)
	}

	status_code := resp.StatusCode
	statusStr := strconv.Itoa(status_code)
	if strings.HasPrefix(statusStr, "4") {
		log.Print("Status Code : ", statusStr)
		return

	} else {
		fmt.Println(status_code, ":", file_Link_format)
		log.Fatal("unable to connect the link : ", file_Link_format)

	}

}

func download_file() {
	full_file_path := filepath.Join(BSE_file_path, file_name)
	if file_exist(full_file_path) {
		fmt.Println("File already exists.")
		return
	}
	URL_status_check()
	resp, err := grab.Get(BSE_file_path, file_Link_format)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)
}
func main() {
	download_file()

}
