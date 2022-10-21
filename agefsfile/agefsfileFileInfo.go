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
	name    string      // base name of the file
	size    int64       // length in bytes for regular files; system-dependent for others
	mode    os.FileMode // file mode bits
	modTime time.Time   // modification time
	isDir   bool        // abbreviation for Mode().IsDir()
	sys     any         // underlying data source (can return nil)
}

// IsDir implements fs.FileInfo
func (a *agefsfileFileInfo) IsDir() bool {
	return a.isDir
}

// ModTime implements fs.FileInfo
func (a *agefsfileFileInfo) ModTime() time.Time {
	return a.modTime
}

// Mode implements fs.FileInfo
func (a *agefsfileFileInfo) Mode() fs.FileMode {
	return a.mode
}

// Name implements fs.FileInfo
func (a *agefsfileFileInfo) Name() string {
	return a.name
}

// Size implements fs.FileInfo
func (a *agefsfileFileInfo) Size() int64 {
	return a.size
}

// Sys implements fs.FileInfo
func (a *agefsfileFileInfo) Sys() any {
	return a.sys
}
