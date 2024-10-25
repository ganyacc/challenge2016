package models

import (
	"encoding/csv"
	"io"
	"os"
)

func LoadLocationsFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// CSV format: City Code, Province Code, Country Code, City Name, Province Name, Country Name
		cityCode := record[0]
		provinceCode := record[1]
		countryCode := record[2]
		cityName := record[3]
		provinceName := record[4]
		countryName := record[5]

		GlobalLocationDB.AddLocation(
			cityCode,
			provinceCode,
			countryCode,
			cityName,
			provinceName,
			countryName,
		)
	}

	return nil
}
