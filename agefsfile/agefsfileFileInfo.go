package agefsfile

import (
	"io/fs"
	"os"
	"time"
)

// Interface guards
var (
	_ os.FileInfo = (*agefsfileFileInfo)(nil)
)

type agefsfileFileInfo struct {
	fi      os.FileInfo
	newname string
	size    int64
}

func NewFileInfo(fi os.FileInfo, newname string, size int64) agefsfileFileInfo {
	return agefsfileFileInfo{
		fi:      fi,
		newname: newname,
		size:    size,
	}
}
func (a *agefsfileFileInfo) FileInfo() os.FileInfo {
	return a
}

// IsDir implements fs.FileInfo
func (a *agefsfileFileInfo) IsDir() bool {
	return a.fi.IsDir()
}

// ModTime implements fs.FileInfo
func (a *agefsfileFileInfo) ModTime() time.Time {
	return a.fi.ModTime()
}

// Mode implements fs.FileInfo
func (a *agefsfileFileInfo) Mode() fs.FileMode {
	return a.fi.Mode()
}

// Name implements fs.FileInfo
func (a *agefsfileFileInfo) Name() string {
	return a.newname
}

// Size implements fs.FileInfo
func (a *agefsfileFileInfo) Size() int64 {
	return a.size
}

// Sys implements fs.FileInfo
func (a *agefsfileFileInfo) Sys() any {
	return a.fi.Sys()
}
