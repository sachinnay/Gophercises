package vault

import (
	"crypto/cipher"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/dash/dash_utils/dashtest"

	homedir "github.com/mitchellh/go-homedir"
)

func TestFile(t *testing.T) {
	v := File("sampleKey", "Samplefile")

	assert.Equalf(t, "sampleKey", v.encodingKey, "Should be equal")
}
func getPath() string {
	home, _ := homedir.Dir()
	filePath := filepath.Join(home, ".testSecret")
	return filePath
}
func removeFile() {
	home, _ := homedir.Dir()
	filePath := filepath.Join(home, ".testSecret")
	os.Remove(filePath)
}

var v = Vault{}

func TestLoad_FilePathError(t *testing.T) {
	err := v.load()
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}

}
func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestLoad_Error(t *testing.T) {

	v := File("???", getPath())
	tempFun := DecryptReaderVar
	defer func() {
		DecryptReaderVar = tempFun
	}()
	DecryptReaderVar = func(key string, r io.Reader) (*cipher.StreamReader, error) {
		return nil, errors.New("Custmised Error")
	}

	err := v.load()
	assert.Errorf(t, err, "Custmised Error", "Should be equal")

}
func TestSave_Error(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, ".testSecret")
	v := File("???", fp)
	tempFun := EncryptWriterVar
	defer func() {
		EncryptWriterVar = tempFun
	}()
	EncryptWriterVar = func(key string, w io.Writer) (*cipher.StreamWriter, error) {
		return nil, errors.New("Custmised Error")
	}

	err := v.save()
	assert.Errorf(t, err, "Custmised Error", "Should be equal")

}

func TestSave_FilePathError(t *testing.T) {
	err := v.save()
	assert.Errorf(t, err, "open : no such file or directory", "Should be equal")

}

func TestSet(t *testing.T) {
	//removeFile()
	v := File("key", getPath())
	key := "twitter_api_key"
	err := v.Set(key, "some-value")
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestSet_Error(t *testing.T) {
	v := File("", getPath())
	err := v.Set("xyz", "testing")
	if err == nil {
		t.Error("Expected  Error but got sucessful execution")
	}
}
func TestSet_Error2(t *testing.T) {
	v := File("", getPath())
	err := v.Set("?", "testing")
	if err == nil {
		t.Error("Expected  Error but got sucessful execution")
	}
}
func TestGet(t *testing.T) {
	//removeFile()
	v := File("key", getPath())
	key := "twitter_api_key"
	value, err := v.Get(key)
	assert.Equalf(t, "some-value", value, "Should be equal")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGet_Error(t *testing.T) {
	//removeFile()
	v := File("key", getPath())
	key := "twitter_api_"
	value, err := v.Get(key)
	if !strings.Contains(err.Error(), "no value for given key") {
		t.Errorf(err.Error(), value)
	}
}
func TestGet_Error2(t *testing.T) {
	//removeFile()
	v := File("", getPath())
	key := "twitter_api_"
	_, err := v.Get(key)
	if err == nil {
		t.Error("Expected Error but got NO error ")
	}
}
