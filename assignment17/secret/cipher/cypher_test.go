package cipher

import (
	"crypto/aes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/CloudBroker/dash_utils/dashtest"
)

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
func TestEncryptWriter(t *testing.T) {
	home, _ := homedir.Dir()
	path := filepath.Join(home, ".test_secret")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	cipherstream, err := EncryptWriter("test-key", f)
	if err != nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
	f.Close()
}
func TestEncryptWriterNegative(t *testing.T) {
	testFile, err := os.OpenFile("test", os.O_RDWR|os.O_CREATE, 0777)
	cipherstream, err := EncryptWriter("test-key", testFile)
	if err != nil {
		t.Errorf("Error for %s: %d.", cipherstream, err)
	}
	testFile.Close()
}
func TestDecryptReader(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, ".test_secret")
	f, _ := os.Open(fp)
	defer f.Close()
	_, err := DecryptReader("test-key", f)
	if err != nil {
		t.Errorf("Expected NO error but got following error : %v ", err)
	}
}
func TestDecryptReaderNeg(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, ".test_secret1")
	f, _ := os.Open(fp)
	defer f.Close()
	_, err := DecryptReader("test-key", f)
	if err == nil {
		t.Errorf("Expected NO error but got following error : %v ", err)
	}
}
func TestCheckIV(t *testing.T) {
	iv := make([]byte, aes.BlockSize)
	err := checkIV(10, iv, errors.New("test"))
	if err == nil {
		t.Error("Expected error but got no error")
	}
}
