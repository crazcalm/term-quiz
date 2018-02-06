package quiz

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestFileArgs(t *testing.T) {
	//Need to check
	// 1 - Files have been passed in
	// 2 - files exist

	file1 := filepath.Join("test_data", "abcd.csv")
	file2 := filepath.Join("test_data", "true_false.csv")
	file3 := filepath.Join("test_data", "fill_in_the_blank.csv")
	notAFile := filepath.Join("test_data", "notAFile")
	emptyString := ""

	error1 := "No files were passed in"
	error2 := "empty string cannot be used as a file"
	error3 := "test_data/notAFile does not exist"

	tests := []struct {
		Files         []string
		ExpectedError bool
		Error         string
	}{
		{[]string{}, true, error1},
		{[]string{emptyString}, true, error2},
		{[]string{notAFile}, true, error3},
		{[]string{file1}, false, ""},
		{[]string{file1, file2}, false, ""},
		{[]string{file1, file2, file3}, false, ""},
		{[]string{file1, emptyString, file2}, true, error2},
		{[]string{file1, notAFile, file2}, true, error3},
	}

	for _, test := range tests {
		err := FileArgs(test.Files)

		if test.ExpectedError && err == nil {
			t.Error("Was expecting an error but no error was received")
		}

		if !test.ExpectedError && err != nil {
			t.Errorf("Unexpected Error: %s", err.Error())
		}

		if test.ExpectedError && err != nil {
			if !strings.Contains(err.Error(), test.Error) {
				t.Errorf("Expected error to contain %s, but the error was actually '%s'", test.Error, err.Error())
			} else {
				continue
			}
		}
	}
}
