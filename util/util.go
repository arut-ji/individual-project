package util

import "encoding/base64"

func DecodeContent(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
