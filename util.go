package quiz

import (
	"fmt"
	"os"
	"strings"
)

//FileArgs -- Checks the file args passed into the program
func FileArgs(files []string) (err error) {
	//At least one file must exist
	if len(files) == 0 {
		return fmt.Errorf("No files were passed in")
	}

	//Make sure that each files exists
	//If it does exist, make sure it is not a directory
	for _, file := range files {
		if strings.EqualFold(file, "") {
			return fmt.Errorf("empty string cannot be used as a file")
		}

		if f, err := os.Stat(file); err != nil {
			if os.IsNotExist(err) {
				//return err
				return fmt.Errorf("%s does not exist", file)
			} else if f.IsDir() {
				return fmt.Errorf("%s is not a file", f.Name())
			}
		}
	}
	return nil
}
