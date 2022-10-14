package file

import (
	"io/fs"
	"os"
	"path"

	"github.com/dhcgn/age-fs/fileinfo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/webdav"
)

func NewFile(logger *logrus.Entry, fi fs.FileInfo, path string) webdav.File {
	osFile, err := os.Open(path)
	if err != nil {
		logger.Error(err)
	}

	return file{
		Logger:   logger,
		FileInfo: fi,
		Path:     path,
		osFile:   osFile,
	}
}

type file struct {
	Logger   *logrus.Entry
	FileInfo fs.FileInfo
	Path     string
	osFile   *os.File
}

// Close implements webdav.File
func (f file) Close() error {
	f.Logger.Debugln("Close")
	return f.osFile.Close()
}

// Read implements webdav.File
func (f file) Read(p []byte) (n int, err error) {
	f.Logger.Debugln("Read")
	return f.osFile.Read(p)
}

// Seek implements webdav.File
func (f file) Seek(offset int64, whence int) (int64, error) {
	f.Logger.WithField("offset", offset).WithField("whence", whence).Debugln("Seek")

	return f.osFile.Seek(offset, whence)
}

// Readdir implements webdav.File
func (f file) Readdir(count int) ([]fs.FileInfo, error) {
	f.Logger.Debugln("Readdir", count)

	dir, err := os.ReadDir(f.Path)
	if err != nil {
		return nil, err
	}

	var fis []fs.FileInfo
	for _, fi := range dir {
		p, err := os.Stat(path.Join(f.Path, fi.Name()))
		if err != nil {
			continue
		}

		fis = append(fis, fileinfo.NewFileInfo(f.Logger, p))
	}

	return fis, nil
}

// Stat implements webdav.File
func (f file) Stat() (fs.FileInfo, error) {
	f.Logger.Debugln("Stat")
	return f.FileInfo, nil
}

// Write implements webdav.File
func (f file) Write(p []byte) (n int, err error) {
	f.Logger.Debugln("Write", len(p), "bytes")

	return f.osFile.Write(p)
}
