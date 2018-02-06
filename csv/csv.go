package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//Read -- Parses csv file
func Read(path string, records [][]string) ([][]string, error) {
	//Make sure the file exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = fmt.Errorf("file: %s does not exist", path)
		return records, err
	}

	//Open the file
	file, err := os.Open(path)
	if err != nil {
		return records, err
	}
	defer file.Close() // nolint: errcheck

	//Read the file
	r := csv.NewReader(file)
	r.LazyQuotes = true // Needed to except the existence of quotes within the statement fields

	//counter used to skip the first record
	count := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return records, fmt.Errorf("Err while reading file %s: %s", file.Name(), err.Error())
		}
		if count != 0 {
			records = append(records, record)
		}
		count++
	}

	return records, nil
}
