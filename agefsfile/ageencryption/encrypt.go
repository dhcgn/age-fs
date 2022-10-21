package ageencryption

import (
	"io"

	"filippo.io/age"
)

func Encrypt(pt io.WriterTo, ct io.Writer, r age.Recipient) error {
	wc, err := age.Encrypt(ct, r)
	if err != nil {
		return err
	}

	pt.WriteTo(wc)

	return err
}
