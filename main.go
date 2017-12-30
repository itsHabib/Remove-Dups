package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Defined as global variables to be easily used in the walk function
// regex string dependent on -dup and -all flags
var fileRegexString string

// used to determine which filename duplicate is deleted if  -dup=1
// is indicated only filenames with example(1) be deleted, and etc
var minusDup *int

// used to delete all duplicates example(1)...example(n), ignores the -dup flag
var minusAll *bool

// traverses through each file in the path given deleting files based
// on flags given
func walkAndDelete(path string, f os.FileInfo, err error) error {
	// depending on whether -all or -dup is used, the regex string is different
	// one to search for all numbers in parentheses or one to search with value
	// of dup flag
	if *minusAll {
		fileRegexString = "\\w+\\([1-9]\\d*\\)($|\\.\\w+$)"
	} else {
		fileRegexString = fmt.Sprintf("\\w+\\(%d\\)($|\\.\\w+$)", *minusDup)
	}
	regex, err := regexp.Compile(fileRegexString)
	if err != nil {
		fmt.Printf("Error in Regex String:%s\n", fileRegexString)
		return err
	}

	if path == "." {
		return nil
	}
	fileName := filepath.Base(path)
	// delete file if regex matches filename
	if regex.MatchString(fileName) {
		err = os.Remove(fileName)
		if err != nil {
			fmt.Printf("Error deleting %s", fileName)
			return err
		}
	}
	return nil
}
