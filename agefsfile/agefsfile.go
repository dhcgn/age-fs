package agefsfile

import (
	"bytes"
	"io/fs"
	"os"

	"filippo.io/age"
	"github.com/dhcgn/age-fs/agefsfile/ageencryption"
	"golang.org/x/net/webdav"
)

// Interface guards
var (
	_ webdav.File = (*agefsfile)(nil)
)

type agefsfile struct {
	PlainContent      []byte
	PlainBuffer       *bytes.Buffer
	PlainReader       *bytes.Reader
	EncryptedFilePath string
	id                *age.X25519Identity
	info              agefsfileFileInfo
}

// Close implements webdav.File
func (af *agefsfile) Close() error {
	f, err := os.OpenFile(af.EncryptedFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	ageencryption.Encrypt(af.PlainReader, f, af.id.Recipient())

	return err
}

// Read implements webdav.File
func (af *agefsfile) Read(p []byte) (n int, err error) {
	return af.PlainReader.Read(p)
}

// Seek implements webdav.File
func (af *agefsfile) Seek(offset int64, whence int) (int64, error) {
	return af.PlainReader.Seek(offset, whence)
}

// Readdir implements webdav.File
func (agefsfile) Readdir(count int) ([]fs.FileInfo, error) {
	panic("Readdir Unimplemented")
}

// Stat implements webdav.File
func (a *agefsfile) Stat() (fs.FileInfo, error) {
	a.info.size = int64(a.PlainBuffer.Len())
	return &a.info, nil
}

// Write implements webdav.File
func (af *agefsfile) Write(p []byte) (n int, err error) {
	n, err = af.PlainBuffer.Write(p)
	if err != nil {
		return 0, err
	}
	af.PlainReader = bytes.NewReader(af.PlainBuffer.Bytes())
	return n, err
}

func New(name string, flag int, perm os.FileMode, id *age.X25519Identity) (*agefsfile, error) {
	path := name + ".age"

	af := &agefsfile{
		EncryptedFilePath: path,
		id:                id,
	}

	// Encrypted existing file
	if stats, err := os.Stat(path); err == nil {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		pt, err := ageencryption.Decrypt(f, id)
		if err != nil {
			return nil, err
		}

		af.PlainBuffer.Write(pt)
		af.info = NewFileInfo(stats, name, int64(af.PlainBuffer.Len()))
	} else {
		af.PlainBuffer = bytes.NewBuffer([]byte{})
	}

	af.PlainReader = bytes.NewReader(af.PlainBuffer.Bytes())

	return af, nil
}
