package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/alecthomas/kingpin"
)

func encrypt(pk *rsa.PublicKey, text, label []byte) ([]byte, error) {
	return rsa.EncryptOAEP(md5.New(), rand.Reader, pk, text, label)
}

var encryptCommand = kingpin.Command("encrypt", "encrypts values for the .travis.yml").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug()
	repo, _, err := client.Repositories.GetFromSlug(s)
	if err != nil {
		return err
	}

	req, err := client.NewRequest("GET", "/repos/"+fmt.Sprint(repo.Id)+"/key", nil, nil)
	if err != nil {
		return err
	}

	var pubkey struct {
		Key         string `json:"key"`
		FingerPrint string `json:"fingerprint"`
	}
	_, err = client.Do(req, &pubkey)
	if err != nil {
		return err
	}

	block, _ := pem.Decode([]byte(pubkey.Key))
	if err != nil {
		return err
	}

	cert, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	rsaPublicKey := cert.(*rsa.PublicKey)

	b, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte("foo=bar"))
	if err != nil {
		return err
	}

	s = "  secure: " + base64.StdEncoding.EncodeToString(b)
	fmt.Println(s)
	return nil
})
var encryptArg = encryptCommand.Arg("data", "data to encrypt").String()
