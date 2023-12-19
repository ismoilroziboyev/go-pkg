package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func HashMD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func HashSHA1(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashSHA256(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashSHA512(s string) string {
	hasher := sha512.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
