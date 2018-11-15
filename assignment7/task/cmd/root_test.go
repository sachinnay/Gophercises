package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.ibm.com/dash/dash_utils/dashtest"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sachinnay/Gophercises/assignment7/task/db"
	"github.com/stretchr/testify/assert"
)

func setDB() string {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "test.db")
	return dbPath
}
func removeDBFile() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)
}

func TestAddCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"Task123"}
	addCmd.Run(addCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	val := strings.Contains(output, "added")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}

func TestListCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record

	listCmd.Run(listCmd, nil)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "1 Task123")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}
func TestDoCommand_ParseError(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"a"}

	doCmd.Run(doCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Error while parsing inputs")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}
func TestDoCommand_DBError(t *testing.T) {

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"1"}

	doCmd.Run(doCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "database not open")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}
func TestDoCommand_InvalidTaskError(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)
	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"10"}

	doCmd.Run(doCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Invalid task number")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}
func TestDoCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"1"}

	doCmd.Run(doCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Marked task :: \"1\"  as completed.")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}

func TestAddCommand_DbError(t *testing.T) {

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"Task123"}
	addCmd.Run(addCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	val := strings.Contains(output, "database not open")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()

}
func TestListCommand_Notask(t *testing.T) {
	removeDBFile()
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	//a := []string{"Task123"}
	listCmd.Run(listCmd, nil)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}

	output := string(content)
	val := strings.Contains(output, "No tasks are available in list")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Db.Close()
}

func TestListCommand_DbError(t *testing.T) {
	removeDBFile()
	dbPath := setDB()
	db.Init(dbPath)
	db.Db.Close()

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	//a := []string{"Task123"}
	listCmd.Run(listCmd, nil)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}

	output := string(content)
	val := strings.Contains(output, "database not open")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()

}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
