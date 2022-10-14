package agefs

import (
	"io/fs"
	"time"

	"github.com/sirupsen/logrus"
)

type fileinfo struct {
	Logger *logrus.Entry
	isDir  bool
}

// IsDir implements fs.FileInfo
func (fi fileinfo) IsDir() bool {
	fi.Logger.Debugln("IsDir", fi.isDir)
	return fi.isDir
}

// ModTime implements fs.FileInfo
func (fi fileinfo) ModTime() time.Time {
	fi.Logger.Debugln("ModTime")
	return time.Now()
}

// Mode implements fs.FileInfo
func (fi fileinfo) Mode() fs.FileMode {
	fi.Logger.Debugln("Mode")
	return fs.FileMode(0)
}

// Name implements fs.FileInfo
func (fi fileinfo) Name() string {
	fi.Logger.Debugln("Name")
	return "testfile"
}

// Size implements fs.FileInfo
func (fi fileinfo) Size() int64 {
	fi.Logger.Debugln("Size")
	return int64(len("Hello"))
}

// Sys implements fs.FileInfo
func (fi fileinfo) Sys() any {
	fi.Logger.Debugln("Sys")
	return nil
}
