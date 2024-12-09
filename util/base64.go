package util

import "encoding/base64"

func Encode_Base64(auth string) string {
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(len(auth)))
	enc.Encode(encoded, []byte(auth))
	return string(encoded)
}

func Decode_Base64(auth string) ([]byte, error) {
	enc := base64.StdEncoding
	decoded := make([]byte, enc.DecodedLen(len(auth)))
	n, err := enc.Decode(decoded, []byte(auth))
	return decoded[:n], err
}
