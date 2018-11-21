package cobra

import (
	"crypto/cipher"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/sachinnay/Gophercises/assignment17/secret/vault"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestSet(t *testing.T) {
	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	args := []string{"test_api_key", "test_api_value"}
	setCmd.Run(setCmd, args)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Value sets successfully")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()

}

func TestSet_Error(t *testing.T) {
	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	args := []string{"-k", "123", "test_api_key1", "test_api_value1"}
	tempFun := vault.DecryptReaderVar
	defer func() {
		vault.DecryptReaderVar = tempFun
	}()
	vault.DecryptReaderVar = func(key string, r io.Reader) (*cipher.StreamReader, error) {
		return nil, errors.New("Custmised Error")
	}

	setCmd.Run(setCmd, args)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Custmised Error")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()

}
func TestGet(t *testing.T) {
	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	args := []string{"test_api_key"}
	getCmd.Run(getCmd, args)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	val := strings.Contains(output, "test_api_key=test_api_value")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()

}
func TestGet_NoValueError(t *testing.T) {
	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	args := []string{"sample11"}
	getCmd.Run(getCmd, args)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	val := strings.Contains(output, "No value set")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	os.Remove("")

}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
