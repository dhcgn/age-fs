package ageencryption

import (
	"io"

	"filippo.io/age"
)

func Decrypt(ct io.Reader, id *age.X25519Identity) ([]byte, error) {
	r, err := age.Decrypt(ct, id)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}
