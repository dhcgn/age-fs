package agefs

import (
	"io"
	"io/fs"

	"github.com/sirupsen/logrus"
)

type file struct {
	Logger   *logrus.Entry
	FileInfo fs.FileInfo
}

// Close implements webdav.File
func (f file) Close() error {
	f.Logger.Debugln("Close")
	return nil
}

// Read implements webdav.File
func (f file) Read(p []byte) (n int, err error) {
	f.Logger.Debugln("Read")
	data := []byte("Hello")
	copy(p, data)
	return len(data), io.EOF
}

// Seek implements webdav.File
func (f file) Seek(offset int64, whence int) (int64, error) {
	f.Logger.Debug("Seek", offset, whence)
	return 0, nil
}

// Readdir implements webdav.File
func (f file) Readdir(count int) ([]fs.FileInfo, error) {
	f.Logger.Debugln("Readdir", count)
	return []fs.FileInfo{

		fileinfo{
			Logger: f.Logger,
		},
	}, nil
}

// Stat implements webdav.File
func (f file) Stat() (fs.FileInfo, error) {
	f.Logger.Debugln("Stat")
	return f.FileInfo, nil
}

// Write implements webdav.File
func (f file) Write(p []byte) (n int, err error) {
	f.Logger.Debugln("Write", len(p), "bytes")
	return 0, nil
}
