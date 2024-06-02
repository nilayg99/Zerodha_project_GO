package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

var (
	currentDate    = time.Now().Format("020106")
	yesterdayDate  = time.Now().AddDate(0, 0, -1).Format("020106")
	fileLinkFormat = "https://www.bseindia.com/download/BhavCopy/Equity/EQ" + yesterdayDate + "_CSV.ZIP"
	zipFileName    = "EQ" + yesterdayDate + "_CSV.ZIP"
	bseFilePath    = "D:/Project/Project_New/Zerodha_project/Zerodha_project_GO/BSE_File_Saved"
)

// fileExists checks if the file exists at the given path.
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// checkURLStatus checks the status of the URL.
func checkURLStatus() {
	resp, err := http.Get(fileLinkFormat)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	statusStr := strconv.Itoa(statusCode)
	if strings.HasPrefix(statusStr, "4") {
		log.Print("Status Code: ", statusStr)
		return
	}
	fmt.Println(statusCode, ":", fileLinkFormat)
	log.Fatal("Unable to connect to the link: ", fileLinkFormat)
}

// extractZIP calls the Python script to extract the ZIP file.
func extractZIP() {
	zipExtractor := "D:/Project/Project_New/Zerodha_project/Zerodha_project_GO/Zip_extractor.py"
	cmd := exec.Command("python", zipExtractor)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running the script:", err)
		fmt.Println("Output:", string(output))
		return
	}
	fmt.Println("ZIP extraction output:", string(output))
}

// downloadFile downloads the ZIP file if it doesn't already exist.
func downloadFile() {
	fullFilePath := filepath.Join(bseFilePath, zipFileName)
	if fileExists(fullFilePath) {
		fmt.Println("File already exists.")
		extractZIP()
		return
	}

	checkURLStatus()

	resp, err := grab.Get(bseFilePath, fileLinkFormat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)
	extractZIP()
}

// timeCheck calls the Python script to check the time.
func timeCheck() {
	timeCheckScript := "D:/Project/Project_New/Zerodha_project/Zerodha_project_GO/Time.py"
	cmd := exec.Command("python", timeCheckScript)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running the script:", err)
		fmt.Println("Output:", string(output))
		return
	}

	result := strings.TrimSpace(string(output))

	if result == "BSE closed no data found !" {
		fmt.Println("No Data found!")
	} else {
		downloadFile()
	}
}

func main() {
	timeCheck()
}
