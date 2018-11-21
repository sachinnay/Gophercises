package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

//This file conatins cipher encryption  & decryption code

//EncryptWriter takes key and io writer as input and decrypt the data and return text in stream writer
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {

	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)
	stream, _ := encrytStream(key, iv)
	n, err := w.Write(iv)
	err = checkIV(n, iv, err)
	return &cipher.StreamWriter{S: stream, W: w}, err
}

//checks the IV and write error
func checkIV(n int, iv []byte, err error) error {
	if len(iv) != n || err != nil {
		return errors.New("Unable to write IV into writer")
	}
	return nil
}

//DecryptReader takes key and io reader as input.
//reader will be carring the encrypted data which will be decrypted and returned in stream writer as plain text
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to read the full iv")
	}
	stream, err := decryptStream(key, iv)
	return &cipher.StreamReader{S: stream, R: r}, err
}

//Used to encrypt the stream
func encrytStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newBlockCipher(key)
	return cipher.NewCFBEncrypter(block, iv), err
}

//Used to decrypt the stream
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newBlockCipher(key)
	return cipher.NewCFBDecrypter(block, iv), err
}

// Create new cipher from key
func newBlockCipher(key string) (cipher.Block, error) {

	hash := md5.New()
	fmt.Fprint(hash, key)
	cipherKey := hash.Sum(nil)
	return aes.NewCipher(cipherKey)
}
