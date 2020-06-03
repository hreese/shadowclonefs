package shadowclonefs

// this is heavily inspired by http://marcio.io/2015/07/calculating-multiple-file-hashes-in-a-single-pass/

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

type MultiHash struct {
	MD5    string `json:"md5"`
	SHA1   string `json:"sha1"`
	SHA256 string `json:"sha256"`
	SHA512 string `json:"sha512"`
}

func CalculateHashes(r io.Reader) (MultiHash, error) {
	var (
		err        error = nil
		md5hash          = md5.New()
		sha1hash         = sha1.New()
		sha256hash       = sha256.New()
		sha512hash       = sha512.New()
	)

	pagesize := os.Getpagesize()

	// read input stream in chunks of the OS's page size
	reader := bufio.NewReaderSize(r, pagesize)

	// create Writer that copies to all hashes
	multiWriter := io.MultiWriter(md5hash, sha1hash, sha256hash, sha512hash)

	// copy the input to all hashes
	_, err = io.Copy(multiWriter, reader)

	return MultiHash{
		hex.EncodeToString(md5hash.Sum(nil)),
		hex.EncodeToString(sha1hash.Sum(nil)),
		hex.EncodeToString(sha256hash.Sum(nil)),
		hex.EncodeToString(sha512hash.Sum(nil)),
	}, err
}
