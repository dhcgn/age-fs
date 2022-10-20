package agefsfile

import (
	"io/fs"

	"filippo.io/age"
)

type agefsfile struct{}

// Close implements webdav.File
func (agefsfile) Close() error {
	panic("unimplemented")
}

// Read implements webdav.File
func (agefsfile) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// Seek implements webdav.File
func (agefsfile) Seek(offset int64, whence int) (int64, error) {
	panic("unimplemented")
}

// Readdir implements webdav.File
func (agefsfile) Readdir(count int) ([]fs.FileInfo, error) {
	panic("unimplemented")
}

// Stat implements webdav.File
func (agefsfile) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}

// Write implements webdav.File
func (agefsfile) Write(p []byte) (n int, err error) {
	panic("unimplemented")
}

func New(name string, id *age.X25519Identity) agefsfile {
	return agefsfile{}
}
