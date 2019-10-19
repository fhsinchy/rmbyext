package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
)

func main() {
	args := os.Args[1:] // getting the arguments except the filename

	if len(args) > 0 {
		p, err := os.Getwd() // getting the working directory
		if err != nil {
			fmt.Println(err)
			return
		}

		var deletedFileCount int
		var log string

		for _, arg := range args {
			var files []string = createFileList(p, arg) // creating file slice for each argument

			if len(files) > 0 {
				removeFiles(arg, files, &deletedFileCount, &log) // calling remove file function on the file list
			} else {
				fmt.Printf("NO %v FILES TO REMOVE.\n", strings.ToUpper(arg))
			}
		}

		if deletedFileCount > 0 {
			writeLogToFile(&log, &deletedFileCount) // writing delete log to a log file
		}
	}
}

func createFileList(p string, arg string) []string {
	var files []string

	filepath.Walk(p, func(path string, file os.FileInfo, _ error) error {
		if !file.IsDir() {
			if filepath.Ext(path) == "." + arg {
				absolutePath, err := filepath.Abs(path) // returning the absolute file to current file if it matches the passed extension
				if err != nil {
					fmt.Println(err)
				}
				files = append(files, absolutePath) // appending the absolute path to our slice
			}
		}
		return nil
	})
	return files // returning the files slice to caller
}

func removeFiles(arg string, files []string, deletedFileCount *int, log *string) {
	fmt.Printf("REMOVING ALL '%v' FILES -->\n", strings.ToUpper(arg))
	for _, file := range files {
		err := os.Remove(file) // removing each file in the slice
		if err != nil {
			fmt.Println(err)
			continue
		}
		*deletedFileCount++ // increasing the number of deleted files
		fmt.Println(file) // printing filename on the console
		*log += file + "\n" // appending filename to the log string
	}
	fmt.Print("\n") // printing an extra newline to separate filename of different extension on the console output
	*log += "\n" // appending an extra newline to separate filename of different extension
}

func writeLogToFile(log *string, deletedFileCount *int) {
	*log += fmt.Sprintf("%v FILES DELETED.\n", *deletedFileCount) // appending number of total files deleted at the of the log string
	err := ioutil.WriteFile("delete_log.log", []byte(*log), 0666) // writing the log string to a log file
	if err != nil {
		fmt.Println(err)
		return
	}
}