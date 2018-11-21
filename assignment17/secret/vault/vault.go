package vault

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/pkg/errors"
	"github.com/sachinnay/Gophercises/assignment17/secret/cipher"
)

//This file performs operation like get/set the key values from secret file
//Also does initial load and save operations on same file

//File ::used for setting and returning the Vault object
func File(encodingKey, filePath string) *Vault {
	return &Vault{encodingKey: encodingKey,
		filePath: filePath,
	}
}

//Vault struct contains keyValues and other required info.
type Vault struct {
	encodingKey string
	filePath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

//EncryptWriterVar Taking function as variable to make it coverage
var EncryptWriterVar = cipher.EncryptWriter

//DecryptReaderVar Taking function as variable to make it coverage
var DecryptReaderVar = cipher.DecryptReader

//Reading key values from secret file by decrypting it.
func (v *Vault) load() error {
	file, err := os.Open(v.filePath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer file.Close()

	r, err := DecryptReaderVar(v.encodingKey, file)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)

}

//Write encrypted keyValues to secret file
func (v *Vault) writekeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

//Read keyValues from Reader which is reference to decrypted secret file
func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

//Write key values to secret file by encrypting it.
func (v *Vault) save() error {

	file, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	w, err := EncryptWriterVar(v.encodingKey, file)
	if err != nil {
		return err
	}

	return v.writekeyValues(w)

}

//Get :: load secret file and get value for given key
func (v *Vault) Get(key string) (string, error) {
	err := v.load()
	if err != nil {
		return "", err
	}
	ret, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret:  no value for given key")
	}
	return ret, nil
}

//Set :: load secret file and add new key, value
func (v *Vault) Set(key string, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	return v.save()
}
