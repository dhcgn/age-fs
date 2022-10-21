package agefsfile

import (
	"bytes"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/dhcgn/age-fs/agefsfile/ageencryption"
)

func TestNew_New_CreateNewFile(t *testing.T) {
	// create and open a temporary file
	f, err := os.CreateTemp("", "tmpfile-") // in Go version older than 1.17 you can use ioutil.TempFile
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	encryptedPath := f.Name() + ".age"
	defer os.Remove(encryptedPath)

	af, err := New(f.Name(), os.O_RDWR, 0, ageencryption.TestX25519Identity)
	if err != nil {
		t.Error(err)
	}

	pt := []byte("Hello World!")
	_, err = af.Write(pt)
	if err != nil {
		t.Error(err)
	}

	err = af.Close()
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat(encryptedPath)
	if err != nil {
		t.Error(err)
	}

	encryptedContent, err := os.ReadFile(encryptedPath)
	if err != nil {
		t.Error(err)
	}

	h := encryptedContent[:len(ageencryption.EncryptedFileHeader)]
	if string(h) != ageencryption.EncryptedFileHeader {
		r, _ := regexp.Compile("[a-zA-Z0-9]+")
		h = r.Find(h)
		t.Errorf("encrypted file header is not correct: '%s', excepted %v", h, ageencryption.EncryptedFileHeader)
	}

	decrypted, err := ageencryption.Decrypt(bytes.NewReader(encryptedContent), ageencryption.TestX25519Identity)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(decrypted, pt) {
		t.Errorf(`Expected '%s', got '%s'`, pt, decrypted)
	}
}
