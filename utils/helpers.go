package utils

import (
	"encoding/base64"
	"math/rand"
)

func EncodeURL(url []byte) string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(url)))

	base64.StdEncoding.Encode(dst, url)

	return string(dst)
}

func RandomStringFromEncoding(n int, encoding []rune) string {
	s := make([]rune, n)

	for i := range s {
		s[i] = encoding[rand.Intn(len(encoding))]
	}

	return string(s)
}
