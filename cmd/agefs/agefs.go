package agefs

import "golang.org/x/net/webdav"

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

func NewFS(rootDir string) AgeFS {
	return &agefs{
		fs: filesystem{
			rootDir: rootDir,
		},
	}
}
