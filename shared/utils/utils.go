package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/jxskiss/base62"
	"math/rand"
	"strconv"
)

func DoCryptoAlgo(url string) string {
	randStr := strconv.Itoa(rand.Intn(1000))
	combined := url + randStr
	base62Str := base62.StdEncoding.EncodeToString([]byte(combined))

	hash := sha256.New()
	hashed := hash.Sum([]byte(base62Str))

	base64Str := base64.URLEncoding.EncodeToString(hashed)
	omittedStr := base64Str[:8]
	return omittedStr
}
