package main

import (
	"net/http"

	"golang.org/x/net/webdav"

	"github.com/dhcgn/age-fs/cmd/agefs"
)

func main() {

	a := agefs.NewFS(`c:\temp\agefs`)

	handler := &webdav.Handler{
		FileSystem: a.FS(),
		LockSystem: webdav.NewMemLS(),
	}

	http.ListenAndServe("localhost:8080", handler)
}
