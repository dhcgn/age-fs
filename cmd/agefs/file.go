package agefs

import "io/fs"

type file struct{}

// Close implements webdav.File
func (file) Close() error {
	return nil
}

// Read implements webdav.File
func (file) Read(p []byte) (n int, err error) {
	data := []byte("Hello")
	p = data
	return len("Hello"), nil
}

// Seek implements webdav.File
func (file) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// Readdir implements webdav.File
func (file) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

// Stat implements webdav.File
func (file) Stat() (fs.FileInfo, error) {
	return fi{}, nil
}

// Write implements webdav.File
func (file) Write(p []byte) (n int, err error) {
	return 0, nil
}
