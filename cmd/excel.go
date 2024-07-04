package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Translation represents a translation.
type Translation struct {
	// French is the French translation.
	French string `json:"fr"`
	// Dutch is the Dutch translation.
	Dutch string `json:"nl"`
	// German is the German translation.
	German string `json:"de"`
}

// BasePackage represents the data from the Excel file.
type BasePackage struct {
	// TypeCode is the type code of the vehicle.
	TypeCode string `json:"typeCode"`
	// Maker is the maker of the vehicle.
	Maker string `json:"maker"`
	// Model is the model of the vehicle.
	Model string `json:"model"`
	// Series is the series of the vehicle.
	Series Translation `json:"series"`
	// FuelType is the fuel type of the vehicle.
	FuelType Translation `json:"fuelType"`
}

// Price represents the price
type Price map[string]float64

// Pricing represents the pricing data from the Excel file.
type Pricing struct {
	// Monthly is the monthly pricing of the vehicle.
	Monthly Price `json:"monthly"`
	// Upfront is the upfront pricing of the vehicle.
	Upfront Price `json:"upfront"`
}

// Advanced represents the comfort data from the Excel file.
type Advanced struct {
	BasePackage
	Pricing
}

// Ultimate represents the comfort data from the Excel file.
type Ultimate struct {
	BasePackage
	Pricing
}

// Comfort represents the comfort data from the Excel file.
type Comfort struct {
	BasePackage
	Pricing
}

// Basic represents the basic data from the Excel file.
type Basic struct {
	BasePackage
	Price `json:"price"`
}

// stringToFloat converts a string to a float64.
func stringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func getValueFromRow(row []string, index int) string {
	total := len(row)
	if (total - 1) < index {
		return ""
	}
	return row[index]
}

// getFuelTypeTranslations returns the fuel type translations.
func getFuelTypeTranslations(fuelType string) (nl, fr, de string) {
	switch fuelType {
	case "Diesel":
		nl = "Diesel"
		fr = "Diesel"
		de = "Diesel"
	case "Benzine", "Petrol (unleaded)":
		nl = "Benzine"
		fr = "Essence"
		de = "Benziner"
	case "Hybrid", "Hybrid (Petrol)":
		nl = "Hybride"
		fr = "Hybride"
		de = "Hybrid"
	case "Electric":
		nl = "Elektrisch"
		fr = "Electrique"
		de = "Vollelektrisch"
	}
	return nl, fr, de
}

// getComfort returns the comfort data from the Excel file.
func getComfort(path string) (data []Comfort, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Upfront
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty for upfront")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	data = []Comfort{}
	for _, row := range rows[1:] {
		rowFuelType := row[5]
		fuelTypeNL, fuelTypeFR, fuelTypeDE := getFuelTypeTranslations(rowFuelType)
		data = append(data, Comfort{
			BasePackage: BasePackage{
				TypeCode: row[0],
				Maker:    "BMW",
				Model:    row[1],
				Series: Translation{
					French: row[3],
					Dutch:  row[2],
					German: row[4],
				},
				FuelType: Translation{
					French: fuelTypeFR,
					Dutch:  fuelTypeNL,
					German: fuelTypeDE,
				},
			},
			Pricing: Pricing{
				Upfront: Price{
					"36/40000":  stringToFloat(getValueFromRow(row, 6)),
					"36/60000":  stringToFloat(getValueFromRow(row, 7)),
					"48/60000":  stringToFloat(getValueFromRow(row, 8)),
					"48/80000":  stringToFloat(getValueFromRow(row, 9)),
					"48/120000": stringToFloat(getValueFromRow(row, 10)),
					"60/50000":  stringToFloat(getValueFromRow(row, 11)),
					"60/80000":  stringToFloat(getValueFromRow(row, 12)),
					"60/100000": stringToFloat(getValueFromRow(row, 13)),
					"60/150000": stringToFloat(getValueFromRow(row, 14)),
					"60/200000": stringToFloat(getValueFromRow(row, 15)),
				},
			},
		})
	}

	// Monthly
	sheetName = f.GetSheetName(1)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty for monthly")
	}

	rows, err = f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] {
		// find the data index
		for i, d := range data {
			if d.BasePackage.TypeCode == row[0] {
				data[i].Pricing.Monthly = Price{
					"36/40000":  stringToFloat(getValueFromRow(row, 6)),
					"36/60000":  stringToFloat(getValueFromRow(row, 7)),
					"48/60000":  stringToFloat(getValueFromRow(row, 8)),
					"48/80000":  stringToFloat(getValueFromRow(row, 9)),
					"48/120000": stringToFloat(getValueFromRow(row, 10)),
					"60/50000":  stringToFloat(getValueFromRow(row, 11)),
					"60/80000":  stringToFloat(getValueFromRow(row, 12)),
					"60/100000": stringToFloat(getValueFromRow(row, 13)),
					"60/150000": stringToFloat(getValueFromRow(row, 14)),
					"60/200000": stringToFloat(getValueFromRow(row, 15)),
				}
				break
			}
		}
	}

	return data, nil
}

// getAdvanced returns the comfort data from the Excel file.
func getAdvanced(path string) (data []Advanced, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Upfront
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty for upfront")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	data = []Advanced{}
	for _, row := range rows[1:] {
		rowFuelType := row[5]
		fuelTypeNL, fuelTypeFR, fuelTypeDE := getFuelTypeTranslations(rowFuelType)
		data = append(data, Advanced{
			BasePackage: BasePackage{
				TypeCode: row[0],
				Maker:    "BMW",
				Model:    row[1],
				Series: Translation{
					French: row[3],
					Dutch:  row[2],
					German: row[4],
				},
				FuelType: Translation{
					French: fuelTypeFR,
					Dutch:  fuelTypeNL,
					German: fuelTypeDE,
				},
			},
			Pricing: Pricing{
				Upfront: Price{
					"36/40000":  stringToFloat(getValueFromRow(row, 6)),
					"36/60000":  stringToFloat(getValueFromRow(row, 7)),
					"48/60000":  stringToFloat(getValueFromRow(row, 8)),
					"48/80000":  stringToFloat(getValueFromRow(row, 9)),
					"48/120000": stringToFloat(getValueFromRow(row, 10)),
					"60/50000":  stringToFloat(getValueFromRow(row, 11)),
					"60/80000":  stringToFloat(getValueFromRow(row, 12)),
					"60/100000": stringToFloat(getValueFromRow(row, 13)),
					"60/150000": stringToFloat(getValueFromRow(row, 14)),
					"60/200000": stringToFloat(getValueFromRow(row, 15)),
				},
			},
		})
	}

	// Monthly
	sheetName = f.GetSheetName(1)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty for monthly")
	}

	rows, err = f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] {
		// find the data index
		for i, d := range data {
			if d.BasePackage.TypeCode == row[0] {
				data[i].Pricing.Monthly = Price{
					"36/40000":  stringToFloat(getValueFromRow(row, 6)),
					"36/60000":  stringToFloat(getValueFromRow(row, 7)),
					"48/60000":  stringToFloat(getValueFromRow(row, 8)),
					"48/80000":  stringToFloat(getValueFromRow(row, 9)),
					"48/120000": stringToFloat(getValueFromRow(row, 10)),
					"60/50000":  stringToFloat(getValueFromRow(row, 11)),
					"60/80000":  stringToFloat(getValueFromRow(row, 12)),
					"60/100000": stringToFloat(getValueFromRow(row, 13)),
					"60/150000": stringToFloat(getValueFromRow(row, 14)),
					"60/200000": stringToFloat(getValueFromRow(row, 15)),
				}
				break
			}
		}
	}

	return data, nil
}

// getUltimate returns the comfort data from the Excel file.
func getUltimate(path string) (data []Ultimate, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Upfront
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty for upfront")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	data = []Ultimate{}
	for _, row := range rows[1:] {
		rowFuelType := row[5]
		fuelTypeNL, fuelTypeFR, fuelTypeDE := getFuelTypeTranslations(rowFuelType)
		data = append(data, Ultimate{
			BasePackage: BasePackage{
				TypeCode: row[0],
				Maker:    "BMW",
				Model:    row[1],
				Series: Translation{
					French: row[3],
					Dutch:  row[2],
					German: row[4],
				},
				FuelType: Translation{
					French: fuelTypeFR,
					Dutch:  fuelTypeNL,
					German: fuelTypeDE,
				},
			},
			Pricing: Pricing{
				Upfront: Price{
					"36/40000":  stringToFloat(getValueFromRow(row, 6)),
					"36/60000":  stringToFloat(getValueFromRow(row, 7)),
					"48/60000":  stringToFloat(getValueFromRow(row, 8)),
					"48/80000":  stringToFloat(getValueFromRow(row, 9)),
					"48/120000": stringToFloat(getValueFromRow(row, 10)),
					"60/50000":  stringToFloat(getValueFromRow(row, 11)),
					"60/80000":  stringToFloat(getValueFromRow(row, 12)),
					"60/100000": stringToFloat(getValueFromRow(row, 13)),
					"60/150000": stringToFloat(getValueFromRow(row, 14)),
					"60/200000": stringToFloat(getValueFromRow(row, 15)),
				},
			},
		})
	}

	// Monthly
	sheetName = f.GetSheetName(1)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty for monthly")
	}

	rows, err = f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] {
		// find the data index
		for i, d := range data {
			if d.BasePackage.TypeCode == row[0] {
				data[i].Pricing.Monthly = Price{
					"36/40000":  stringToFloat(getValueFromRow(row, 6)),
					"36/60000":  stringToFloat(getValueFromRow(row, 7)),
					"48/60000":  stringToFloat(getValueFromRow(row, 8)),
					"48/80000":  stringToFloat(getValueFromRow(row, 9)),
					"48/120000": stringToFloat(getValueFromRow(row, 10)),
					"60/50000":  stringToFloat(getValueFromRow(row, 11)),
					"60/80000":  stringToFloat(getValueFromRow(row, 12)),
					"60/100000": stringToFloat(getValueFromRow(row, 13)),
					"60/150000": stringToFloat(getValueFromRow(row, 14)),
					"60/200000": stringToFloat(getValueFromRow(row, 15)),
				}
				break
			}
		}
	}

	return data, nil
}

// getBasic returns the basic data from the Excel file.
func getBasic(path string) (data []Basic, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, errors.New("sheet name is empty")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	data = []Basic{}
	for _, row := range rows[2:] {
		rowFuelType := row[5]
		fuelTypeNL, fuelTypeFR, fuelTypeDE := getFuelTypeTranslations(rowFuelType)
		data = append(data, Basic{
			BasePackage: BasePackage{
				TypeCode: row[0],
				Maker:    "BMW",
				Model:    row[1],
				Series: Translation{
					French: row[3],
					Dutch:  row[2],
					German: row[4],
				},
				FuelType: Translation{
					French: fuelTypeFR,
					Dutch:  fuelTypeNL,
					German: fuelTypeDE,
				},
			},
			Price: Price{
				"36/40000":     stringToFloat(getValueFromRow(row, 6)),
				"36/60000":     stringToFloat(getValueFromRow(row, 7)),
				"48/60000":     stringToFloat(getValueFromRow(row, 8)),
				"48/80000":     stringToFloat(getValueFromRow(row, 9)),
				"60/80000":     stringToFloat(getValueFromRow(row, 10)),
				"60/100000":    stringToFloat(getValueFromRow(row, 11)),
				"72/120000":    stringToFloat(getValueFromRow(row, 12)),
				"48/Unlimited": stringToFloat(getValueFromRow(row, 13)),
				"72/Unlimited": stringToFloat(getValueFromRow(row, 14)),
			},
		})
	}

	return data, nil
}
