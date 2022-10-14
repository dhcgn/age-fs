package agefs

import (
	"github.com/dhcgn/age-fs/filesystem"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/webdav"
)

// FS implements AgeFS
func (agefs *agefs) FS() webdav.FileSystem {
	return agefs.fs
}

type agefs struct {
	fs webdav.FileSystem
}
type AgeFS interface {
	FS() webdav.FileSystem
}

func NewFS(rootDir, privateKey string, log *logrus.Entry) AgeFS {
	return filesystem.NewFileSystem(rootDir, privateKey, log)
}
