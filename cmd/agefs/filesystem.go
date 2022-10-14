package agefs

import (
	"context"
	"io/fs"

	"golang.org/x/net/webdav"
)

type filesystem struct {
	rootDir string
}

// Mkdir implements webdav.FileSystem
func (filesystem) Mkdir(ctx context.Context, name string, perm fs.FileMode) error {
	panic("unimplemented")
}

// OpenFile implements webdav.FileSystem
func (filesystem) OpenFile(ctx context.Context, name string, flag int, perm fs.FileMode) (webdav.File, error) {
	return file{}, nil
}

// RemoveAll implements webdav.FileSystem
func (filesystem) RemoveAll(ctx context.Context, name string) error {
	panic("unimplemented")
}

// Rename implements webdav.FileSystem
func (filesystem) Rename(ctx context.Context, oldName string, newName string) error {
	panic("unimplemented")
}

// Stat implements webdav.FileSystem
func (filesystem) Stat(ctx context.Context, name string) (fs.FileInfo, error) {
	return fi{}, nil
}
