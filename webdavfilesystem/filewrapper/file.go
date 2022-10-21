package filewrapper

import (
	"io/fs"
	"os"
	"strings"

	"golang.org/x/net/webdav"
)

// Interface guards
var (
	_ webdav.File = (*internalfile)(nil)
)

type internalfile struct {
	file *os.File
}

// Close implements webdav.File
func (f *internalfile) Close() error {
	return f.file.Close()
}

// Read implements webdav.File
func (f *internalfile) Read(p []byte) (n int, err error) {
	return f.file.Read(p)
}

// Seek implements webdav.File
func (f *internalfile) Seek(offset int64, whence int) (int64, error) {
	return f.file.Seek(offset, whence)
}

// Readdir implements webdav.File
func (f *internalfile) Readdir(count int) ([]fs.FileInfo, error) {
	fis, err := f.file.Readdir(count)
	if err != nil {
		return nil, err
	}

	var ret []fs.FileInfo
	for _, fi := range fis {
		if strings.HasSuffix(fi.Name(), ".age") && !fi.IsDir() {
			// agefsfileInfo := agefsfile.NewFileInfo(fi, strings.TrimSuffix(fi.Name(), ".age"), fi.Size())
			// ret = append(ret, agefsfileInfo.FileInfo())
			ret = append(ret, fi)
		}

		if fi.IsDir() {
			ret = append(ret, fi)
		}
	}

	return ret, nil
}

// Stat implements webdav.File
func (f *internalfile) Stat() (fs.FileInfo, error) {
	return f.file.Stat()
}

// Write implements webdav.File
func (f *internalfile) Write(p []byte) (n int, err error) {
	return f.file.Write(p)
}

func NewFile(file *os.File) webdav.File {
	return &internalfile{file: file}
}
