package ageencryption

import (
	"bytes"
	_ "embed"
	"io"
	"reflect"
	"testing"

	"filippo.io/age"
)

var (
	//go:embed test_assets/hello.age
	encryptedTestFile []byte
)

func TestDecrypt(t *testing.T) {
	type args struct {
		dec io.Reader
		id  *age.X25519Identity
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Decrypt hello.age",
			args: args{
				dec: bytes.NewReader(encryptedTestFile),
				id:  TestX25519Identity,
			},
			want: []byte("Hello!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.dec, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() = %s, want %s", got, tt.want)
			}
		})
	}
}
