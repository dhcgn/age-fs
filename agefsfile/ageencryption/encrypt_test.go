package ageencryption

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"filippo.io/age"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		pt io.Reader
		r  age.Recipient
	}
	tests := []struct {
		name            string
		args            args
		wantCtSignature string
		wantErr         bool
	}{
		{
			name: "Encrypt hello",
			args: args{
				pt: bytes.NewBufferString("Hello!"),
				r:  Recipient,
			},
			wantCtSignature: "age-encryption.org/v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &bytes.Buffer{}
			if err := Encrypt(tt.args.pt, ct, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCt := ct.String(); !strings.HasPrefix(gotCt, tt.wantCtSignature) {
				t.Errorf("Encrypt() = '%v', want '%v'", gotCt, tt.wantCtSignature)
			}
		})
	}
}
