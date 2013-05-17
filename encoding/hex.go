package encoding

import (
	"encoding/hex"
)

func DecodeHex(raw []byte) ([]byte, error) {
	b, err := hex.DecodeString(string(raw))
	return b, err
}

func EncodeHex(raw []byte) []byte {
	buff := make([]byte, hex.EncodedLen(len(raw)))
	hex.Encode(buff, raw)
	return buff
}
