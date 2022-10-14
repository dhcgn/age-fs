package agefs

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/webdav"
)

// FS implements AgeFS
func (agefs *agefs) FS() webdav.FileSystem {
	return agefs.fs
}

type agefs struct {
	fs filesystem
}
type AgeFS interface {
	FS() webdav.FileSystem
}

func NewFS(rootDir, privateKey string, log *logrus.Entry) AgeFS {
	return &agefs{
		fs: filesystem{
			rootDir:    rootDir,
			privateKey: privateKey,
			Logger:     log,
		},
	}
}
