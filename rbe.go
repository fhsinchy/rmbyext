package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
)

func main() {
	args := os.Args[1:]

	if(len(args) > 0) {
		p, err := os.Getwd()
		
		if(err != nil) {
			panic(err)
		}

		var deletedFileCount int
		var log string

		for _, arg := range args {
			var files []string = createFileList(p, arg)

			if len(files) > 0 {
				removeFiles(arg, files, &deletedFileCount, &log)
			} else {
				fmt.Printf("NO %v FILES TO REMOVE.\n", strings.ToUpper(arg))
			}
		}

		if (deletedFileCount > 0) {
			writeLogToFile(&log, &deletedFileCount)
		}
	}
}

func createFileList(p string, arg string) []string {
	var files []string

	filepath.Walk(p, func(path string, file os.FileInfo, _ error) error {
		if !file.IsDir() {
			if filepath.Ext(path) == "." + arg {
				absolutePath, _ := filepath.Abs(path)
				files = append(files, absolutePath)
			}
		}
		return nil
	})
	return files
}

func removeFiles(arg string, files []string, deletedFileCount *int, log *string) {
	fmt.Printf("REMOVING ALL '%v' FILES -->\n", strings.ToUpper(arg))
	for _, file := range files {
		os.Remove(file)
		*deletedFileCount++
		fmt.Println(file)
		*log += file + "\n"
	}
	*log += "\n"
	fmt.Print("\n")
}

func writeLogToFile(log *string, deletedFileCount *int) {
	*log += fmt.Sprintf("%v FILES DELETED.\n", *deletedFileCount)
	ioutil.WriteFile("delete_log.log", []byte(*log), 0666)
}