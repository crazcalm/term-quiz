package csv

import (
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	abcd := filepath.Join("test_data", "abcd.csv")
	trueFalse := filepath.Join("test_data", "true_false.csv")
	fillInBlank := filepath.Join("test_data", "fill_in_the_blank.csv")
	badData := filepath.Join("test_data", "error.csv")

	var tests = []struct {
		Path        string
		ExpectError bool
		Expected    int
	}{
		{"", true, 0},
		{"notAFile", true, 0},
		{badData, true, 0},
		{abcd, false, 3},
		{trueFalse, false, 3},
		{fillInBlank, false, 3},
		{badData, false, 3},
	}

	for _, test := range tests {
		container := [][]string{}
		result, err := Read(test.Path, container)

		if test.ExpectError && err == nil {
			t.Error("Was expecting an error, but no error was received")
			return
		}

		if !test.ExpectError && err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}

		//I was expecting an error
		//I got an error
		//Test case passed
		if test.ExpectError && err != nil {
			return
		}

		if test.Expected != len(result) {
			t.Errorf("Expected a length of %d, but got %d", test.Expected, len(result))
		}
	}
}
