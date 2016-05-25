package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/alecthomas/kingpin"
)

func toBytes(n *big.Int, significant []byte) []byte {
	n1 := big.NewInt(-1)
	n0 := big.NewInt(0)
	ns := n.String()
	ss := "0"
	if len(significant) > 0 {
		ss = fmt.Sprint(significant[0])
	}

	if n.Cmp(n1) >= 0 && n.Cmp(n0) <= 0 && ((len(ss) >= 7 && len(ns) >= 7 && ss[7] == ns[7]) || (ss == "0" && ns == "0")) {
		if len(significant) == 0 {
			significant = []byte{0}
		}
		return significant
	}
	nn := big.NewInt(0)
	y, m := nn.DivMod(n, big.NewInt(256), big.NewInt(0))
	b := toBytes(y, m.Bytes())
	return append(b, significant...)
}

var pubkeyCommand = kingpin.Command("pubkey", "prints out a repository's public key").Action(func(ctx *kingpin.ParseContext) error {
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
	b1 := toBytes(big.NewInt(int64(rsaPublicKey.E)), nil)
	b2 := toBytes(rsaPublicKey.N, nil)
	buf := new(bytes.Buffer)
	buf.Write([]byte{0, 0, 0, '\a'})
	buf.Write([]byte("ssh-rsa"))
	binary.Write(buf, binary.BigEndian, uint32(len(b1)))
	buf.Write(b1)
	binary.Write(buf, binary.BigEndian, uint32(len(b2)))
	buf.Write(b2)
	s = "ssh-rsa " + base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("Public key for %s:\n\n", s)
	fmt.Println(s)
	return nil
})
var pubkeyRepoFlag = pubkeyCommand.Flag("repo", "repository").Short('r').String()
