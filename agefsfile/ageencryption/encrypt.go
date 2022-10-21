package ageencryption

import (
	"io"

	"filippo.io/age"
)

func Encrypt(pt io.Reader, ct io.Writer, r age.Recipient) error {
	wc, err := age.Encrypt(ct, r)
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = io.Copy(wc, pt)
	if err != nil {
		return err
	}

	return err
}
