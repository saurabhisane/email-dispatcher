package main

import (
	"encoding/csv"
	"os"
)

func loadRecipients(filePath string, ch chan Recipient) error {

	defer close(ch)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		// fmt.Println(record)
		ch <- Recipient{
			Name : record[0],
			Email : record[1],
		}
	}

	return nil

}
