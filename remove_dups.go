package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Defined as global variables to be easily used in the walk function

// used to determine which filename duplicate is deleted if  -dup=1
// is indicated only filenames with example(1) be deleted, and etc
var minusDup *int

// used to delete all duplicates example(1)...example(n), ignores the -dup flag
var minusAll *bool

//used to show files that are being deleted
var minusVerbose *bool

// global variable to count # of files deleted
var numFilesDeleted int

// traverses through each file in the path given deleting files based
// on flags given
func walkAndDelete(path string, f os.FileInfo, err error) error {
	// depending on whether -all or -dup is used, the regex string is different
	// one to search for all numbers in parentheses or one to search with value
	// of dup flag
	// regex string dependent on -dup and -all flags
	var fileRegexString string
	if *minusDup == 0 && !*minusAll {
		fileRegexString = "[\\w\\s]+\\(\\d+\\)($|\\.\\w+$)"
	} else if *minusAll && *minusDup != 0 {
		fileRegexString = "[\\w\\s]+\\(\\d+\\)($|\\.\\w+$)"
	} else if *minusDup != 0 {
		fileRegexString = fmt.Sprintf("[\\w\\s]+\\(%d\\)($|\\.\\w+$)", *minusDup)
	} else {
		fileRegexString = "[\\w\\s]+\\(\\d+\\)($|\\.\\w+$)"
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
		err = os.Remove(path)
		if err != nil {
			fmt.Printf("Error deleting %s\n", fileName)
			return err
		}
		numFilesDeleted++
		if *minusVerbose {
			fmt.Printf("Deleted: %s\n", fileName)
		}
	}
	return nil
}

func main() {
	// Grab values of flags used
	minusDup = flag.Int("dup", 0, "Duplicate Number to Delete")
	minusAll = flag.Bool("all", false, "Delete all duplicate files")
	minusVerbose = flag.Bool("v", false, "Verbose")
	flag.Parse()
	flags := flag.Args()

	// check if path was given
	if len(flags) == 0 {
		fmt.Print("usage: remove_dups <path>")
		os.Exit(1)
	}
	// walk and delete with path given
	Path := flag.Arg(0)
	Path, err := filepath.EvalSymlinks(Path)
	if err != nil {
		fmt.Printf("Error evaluating symlinks %v\n", err)
		os.Exit(1)
	}
	filepath.Walk(Path, walkAndDelete)
	fmt.Printf("%d files deleted\n", numFilesDeleted)
}
