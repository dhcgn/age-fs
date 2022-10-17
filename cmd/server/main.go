package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/webdav"

	"github.com/dhcgn/age-fs/mywebdav"
	"github.com/sirupsen/logrus"
)

const (
	testFolder    = `c:\temp\agefs`
	testPrivatKey = `AGE-SECRET-KEY-1JWDD9FJHFULZMXPJWXVP2WX6J9KU9700HGCA72YH509NQXED6VXS8MLZJ4`
)

var (
	logger = logrus.New()
	log    = logger.WithField("scope", "main")
)

func main() {
	logger.Level = logrus.DebugLevel

	handler := &webdav.Handler{
		FileSystem: mywebdav.NewFileSystem(testFolder),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			log.Debugln(r.Method, r.URL.Path)
			if err != nil {
				log.Error(err)
			}
		},
	}

	go func() {
		//log.Errorln(http.ListenAndServe("localhost:8080", handler))
		log.Fatalln(http.ListenAndServe("localhost:8080", handler))
	}()

	fmt.Println("Server started on port 8080")
	fmt.Println("net use x: http://localhost:8080")

	select {}
}
