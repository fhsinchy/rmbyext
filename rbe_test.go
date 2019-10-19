package main

import(
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	// create temporary files before testing
	for i := 1; i <= 10; i++ {
		os.Remove("file_" + strconv.Itoa(i) + ".extension")
		file, _ := os.Create("file_" + strconv.Itoa(i) + ".extension")

		file.Close()
	}

	// run the tests
	var exitVal int = m.Run()

	// remove temporary files after testing
	for i := 1; i <= 10; i++ {
		os.Remove("file_" + strconv.Itoa(i) + ".extension")
	}
	os.Remove("delete_log.log")

	// exit with the received exit value
	os.Exit(exitVal)
}

func TestCreateFileList(t *testing.T) {
	p, _ := os.Getwd()

	var files []string = createFileList(p, "extension")

	if len(files) != 10 {
		t.Errorf("Expected the number of files to be exactly 10, got %v", len(files))
	}
}

func TestRemoveFilesAndWriteLogToFile(t *testing.T) {
	var deletedFileCount int
	var log string

	p, _ := os.Getwd()

	var files []string = createFileList(p, "extension")

	removeFiles("extension", files, &deletedFileCount, &log)

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

	logFileContent, _ := ioutil.ReadFile("delete_log.log")

	if log != string(logFileContent) {
		t.Error("Expected the log file to have same content as the log variable")
	}
}