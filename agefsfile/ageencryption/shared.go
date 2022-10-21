package ageencryption

import "filippo.io/age"

const (
	pub  = "age100nsfsqrnz5y259jq0f6fjk6wfxchwd0wjsf9cdd06jf5peuedrsg5seqk"
	priv = "AGE-SECRET-KEY-1GHF2H2SGU28TCAQR068AGTSVXY6QP9RJWCQJE66094C5QS3RUDXQTMPD7K"

	EncryptedFileHeader = "age-encryption.org/v1"
)

var (
	TestX25519Identity *age.X25519Identity
	Recipient          age.Recipient
)

func init() {
	id, err := age.ParseX25519Identity(priv)
	if err != nil {
		panic(err)
	}
	TestX25519Identity = id

	Recipient = id.Recipient()
}
