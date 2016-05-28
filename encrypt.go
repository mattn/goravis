package main

import (
	"crypto/rand"
	"crypto/rsa"
	//"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin"
	"github.com/mattn/go-isatty"
)

var encryptCommand = kingpin.Command("encrypt", "encrypts values for the .travis.yml").Action(func(ctx *kingpin.ParseContext) error {
	err := auth()
	if err != nil {
		return err
	}

	s := slug(ctx)
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

	var content []byte
	if len(ctx.Elements) < 2 {
		if runtime.GOOS == "windows" {
			fmt.Println("Reading from stdin, press Ctrl+Z when done")
		} else {
			fmt.Println("Reading from stdin, press Ctrl+D when done")
		}
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
	} else {
		s = ""
		for _, v := range ctx.Elements[1:] {
			if s != "" {
				s += " "
			}
			s += *v.Value
		}
		content = []byte(s)
	}

	b, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, content)
	if err != nil {
		return err
	}

	s = base64.StdEncoding.EncodeToString(b)
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("  secure: " + s)
	} else {
		fmt.Printf("%q", s)
	}
	return nil
})
var encryptArg = encryptCommand.Arg("data", "data to encrypt").Strings()
var encryptRepoFlag = encryptCommand.Flag("repo", "repository").Short('r').String()
