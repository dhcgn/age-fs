package agefs

import (
	"context"
	"io/fs"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/webdav"
)

type filesystem struct {
	rootDir    string
	privateKey string
	Logger     *logrus.Entry
}

// Mkdir implements webdav.FileSystem
func (fs filesystem) Mkdir(ctx context.Context, name string, perm fs.FileMode) error {
	fs.Logger.Debugln("Mkdir", name)
	panic("unimplemented")
}

// OpenFile implements webdav.FileSystem
func (fs filesystem) OpenFile(ctx context.Context, name string, flag int, perm fs.FileMode) (webdav.File, error) {
	fs.Logger.Debugln("OpenFile", name)

	var fi fileinfo
	if name == "/" {
		fi = fileinfo{
			isDir:  true,
			Logger: fs.Logger.WithField("scope", "fileinfo"),
		}
	} else {
		fi = fileinfo{
			isDir:  false,
			Logger: fs.Logger.WithField("scope", "fileinfo"),
		}
	}

	return file{
		Logger:   fs.Logger.WithField("scope", "file"),
		FileInfo: fi,
	}, nil
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

	if name == "/" {
		return fileinfo{
			Logger: fs.Logger.WithField("scope", "fileinfo"),
			isDir:  true,
		}, nil
	}

	return fileinfo{
		Logger: fs.Logger.WithField("scope", "fileinfo"),
	}, nil
}
