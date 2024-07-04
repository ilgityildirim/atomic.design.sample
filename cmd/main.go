package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// EnvCountry is the environment variable for the country
const EnvCountry = "APP_COUNTRY"

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var pathBasic, pathComfort, pathAdvanced, pathUltimate  string
	country := os.Getenv(EnvCountry)
	if country == "" {
		pathBasic = currentDir + "/basic-be.xlsx"
		pathComfort = currentDir + "/comfort-be.xlsx"
		pathAdvanced = currentDir + "/advanced-be.xlsx"
        pathUltimate = currentDir + "/ultimate-be.xlsx"
	} else if country == "be" || country == "lu" {
		pathBasic = fmt.Sprintf("%s/basic-%s.xlsx", currentDir, country)
		pathComfort = fmt.Sprintf("%s/comfort-%s.xlsx", currentDir, country)
		pathAdvanced = fmt.Sprintf("%s/advanced-%s.xlsx", currentDir, country)
        pathUltimate = fmt.Sprintf("%s/ultimate-%s.xlsx", currentDir, country)
	} else {
		log.Fatal("please provide country environment variable: lu or be")
	}

	fmt.Printf("Current Working Directory: %s, Country: %s\n", currentDir, country)
	basic, err := getBasic(pathBasic)
	if err != nil {
		log.Fatal(err)
	}

	comfort, err := getComfort(pathComfort)
	if err != nil {
		log.Fatal(err)
	}

	advanced, err := getAdvanced(pathAdvanced)
    	if err != nil {
    		log.Fatal(err)
    	}

    ultimate, err := getUltimate(pathUltimate)
    	if err != nil {
    		log.Fatal(err)
    	}

	data := struct {
		Basic   []Basic   `json:"basic"`
		Comfort []Comfort `json:"comfort"`
		Advanced []Advanced   `json:"advanced"`
        Ultimate []Ultimate `json:"ultimate"`
	}{
		Basic: basic,
		Comfort: comfort,
		Advanced: advanced,
		Ultimate: ultimate,
	}

	file, _ := json.MarshalIndent(data, "", "  ")
	if err = os.WriteFile("../src/config/data.json", file, 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("data.json created successfully")
}
