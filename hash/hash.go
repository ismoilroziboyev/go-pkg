package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
