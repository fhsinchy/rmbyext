package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

var ext string = "extension"

func TestMain(m *testing.M) {
	// create temporary files before testing
	for i := 1; i <= 10; i++ {
		os.Remove("file_" + strconv.Itoa(i) + "." + ext)
		
		file, err := os.Create("file_" + strconv.Itoa(i) + "." + ext)
		if err != nil {
			fmt.Println(err)
			continue
		}

		file.Close()
	}

	// run the tests
	var exitVal int = m.Run()

	// remove temporary files after testing
	for i := 1; i <= 10; i++ {
		os.Remove("file_" + strconv.Itoa(i) + "." + ext)
	}
	os.Remove("delete_log.log")

	// exit with the received exit value
	os.Exit(exitVal)
}

func TestCreateFileList(t *testing.T) {
	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []string = createFileList(p, ext)

	if len(files) != 10 {
		t.Errorf("Expected the number of files to be exactly 10, got %v", len(files))
	}
}

func TestRemoveFilesAndWriteLogToFile(t *testing.T) {
	var deletedFileCount int
	var log string

	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []string = createFileList(p, ext)

	removeFiles(ext, files, &deletedFileCount, &log)

	if deletedFileCount != 10 {
		t.Errorf("Expected the number of files deleted to be exactly 10, got %v", deletedFileCount)
	}

	if log == "" {
		t.Error("Expected log to have a non zero string value")
	}

	writeLogToFile(&log, &deletedFileCount)

	info, err := os.Stat("delete_log.log")
	
	if os.IsNotExist(err) || info.IsDir() {
		t.Error("Expected a log file to be created")
	}

	logFileContent, err := ioutil.ReadFile("delete_log.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	if log != string(logFileContent) {
		t.Error("Expected the log file to have same content as the log variable")
	}
}