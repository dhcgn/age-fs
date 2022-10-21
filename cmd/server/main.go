package main

import (
	"fmt"
	"net/http"

	"filippo.io/age"
	"golang.org/x/net/webdav"

	"github.com/dhcgn/age-fs/webdavfilesystem"
	"github.com/sirupsen/logrus"
)

const (
	testFolder    = `c:\temp\agefs`
	testPrivatKey = `AGE-SECRET-KEY-102GRTMGVJH4NGJ2JYUPN2TLMFT4CGK33LKAZ5AWF30C25LFAU44SMZLGTN`
)

var (
	logger = logrus.New()
	log    = logger.WithField("scope", "main")
)

func main() {
	logger.Level = logrus.DebugLevel

	i, err := age.ParseX25519Identity(testPrivatKey)
	if err != nil {
		log.Fatal(err)
	}

	handler := &webdav.Handler{
		FileSystem: webdavfilesystem.NewFileSystem(testFolder, i),
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
