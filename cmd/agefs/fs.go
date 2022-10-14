package agefs

import (
	"io/fs"
	"time"
)

type fi struct{}

// IsDir implements fs.FileInfo
func (fi) IsDir() bool {
	return false
}

// ModTime implements fs.FileInfo
func (fi) ModTime() time.Time {
	return time.Now()
}

// Mode implements fs.FileInfo
func (fi) Mode() fs.FileMode {
	return fs.FileMode(0)
}

// Name implements fs.FileInfo
func (fi) Name() string {
	return "testfile"
}

// Size implements fs.FileInfo
func (fi) Size() int64 {
	return int64(len("Hello"))
}

// Sys implements fs.FileInfo
func (fi) Sys() any {
	return nil
}
