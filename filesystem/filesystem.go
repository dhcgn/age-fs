package filesystem

import (
	"context"
	"io/fs"
	"os"
	"path"

	"github.com/dhcgn/age-fs/file"
	"github.com/dhcgn/age-fs/fileinfo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/webdav"
)

func NewFileSystem(rootDir, privateKey string, logger *logrus.Entry) webdav.FileSystem {
	return filesystem{
		rootDir:    rootDir,
		privateKey: privateKey,
		Logger:     logger,
		fsinternal: os.DirFS(rootDir),
	}
}

type filesystem struct {
	rootDir    string
	privateKey string
	Logger     *logrus.Entry
	fsinternal fs.FS
}

// Mkdir implements webdav.FileSystem
func (fs filesystem) Mkdir(ctx context.Context, name string, perm fs.FileMode) error {
	fs.Logger.Debugln("Mkdir", name)
	panic("unimplemented")
}

// OpenFile implements webdav.FileSystem
func (fs filesystem) OpenFile(ctx context.Context, name string, flag int, perm fs.FileMode) (webdav.File, error) {
	fs.Logger.Debugln("OpenFile", name)

	fiMounted, err := fs.Stat(ctx, path.Join(fs.rootDir, name))
	if err != nil {
		return nil, err
	}

	fi := fileinfo.NewFileInfo(fs.Logger.WithField("scope", "fileinfo"), fiMounted)
	return file.NewFile(fs.Logger.WithField("scope", "file"), fi), nil
}

// RemoveAll implements webdav.FileSystem
func (fs filesystem) RemoveAll(ctx context.Context, name string) error {
	fs.Logger.Debugln("RemoveAll", name)
	panic("unimplemented")
}

// Rename implements webdav.FileSystem
func (fs filesystem) Rename(ctx context.Context, oldName string, newName string) error {
	fs.Logger.Debugln("Rename", oldName, newName)
	panic("unimplemented")
}

// Stat implements webdav.FileSystem
func (fs filesystem) Stat(ctx context.Context, name string) (fs.FileInfo, error) {
	fs.Logger.Debugln("Stat", name)

	fiMounted, err := fs.Stat(ctx, path.Join(fs.rootDir, name))
	if err != nil {
		return nil, err
	}
	fi := fileinfo.NewFileInfo(fs.Logger.WithField("scope", "fileinfo"), fiMounted)

	return fi, nil
}
