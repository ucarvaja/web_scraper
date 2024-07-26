package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"web_scraper/internal/config/config"
	"web_scraper/internal/scraper"
)

func main() {
	cfg := config.LoadConfig()

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Enter link")
		fmt.Println("2. Clear JSON entries")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			scraper.EnterLink(cfg.JSONFilePath)
		case 2:
			clearJSON(cfg.JSONFilePath)
		case 3:
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}

func clearJSON(filePath string) {
	// Create an empty JSON file
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}
	defer file.Close()

	// Write an empty JSON object to the file
	emptyData := []scraper.VideoInfo{}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(emptyData)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	fmt.Println("JSON entries cleared.")
}
