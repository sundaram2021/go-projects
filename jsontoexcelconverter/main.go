package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tealeg/xlsx"
)

func CreateJson(data interface{}) {
	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("error in creating file", err)
		return
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error in marshaling data", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("error in writing to file", err)
		return
	}

	fmt.Println("JSON file created successfully")
}

func JsonToExcel() {
	// Open the JSON file
	jsonfile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("error in opening file", err)
		return
	}
	defer jsonfile.Close()

	// Read all the data from the file
	data, err := io.ReadAll(jsonfile)
	if err != nil {
		fmt.Println("error in reading file", err)
		return
	}

	// Unmarshal the data into an interface{}
	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error in unmarshaling the data", err)
		return
	}

	// Create a new Excel file
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("error in creating excel sheet", err)
		return
	}

	// Process the data based on its type
	switch v := result.(type) {
	case []interface{}:
		// Assuming it's an array of objects (like []Person)
		if len(v) > 0 {
			if obj, ok := v[0].(map[string]interface{}); ok {
				// Add headers
				headerRow := sheet.AddRow()
				for key := range obj {
					headerRow.AddCell().Value = key
				}

				// Add rows
				for _, item := range v {
					row := sheet.AddRow()
					if obj, ok := item.(map[string]interface{}); ok {
						for _, value := range obj {
							row.AddCell().Value = fmt.Sprintf("%v", value)
						}
					}
				}
			}
		}
	case map[string]interface{}:
		// Assuming it's a single object (like Person)
		headerRow := sheet.AddRow()
		dataRow := sheet.AddRow()
		for key, value := range v {
			headerRow.AddCell().Value = key
			dataRow.AddCell().Value = fmt.Sprintf("%v", value)
		}
	default:
		fmt.Println("unsupported JSON structure")
		return
	}

	// Save the Excel file
	err = file.Save("data.xlsx")
	if err != nil {
		fmt.Println("error in saving excel file:", err)
		return
	}

	fmt.Println("Excel file successfully created with the given data")
}

func main() {
	// Example JSON data
	data := []map[string]interface{}{
		{"name": "Sundarm", "age": 23, "city": "City A"},
		{"name": "Aman", "age": 26, "city": "City B"},
		{"name": "Dheeraj", "age": 25, "city": "City C"},
		{"name": "Avnish", "age": 21, "city": "City D"},
	}

	CreateJson(data)
	JsonToExcel()
}
