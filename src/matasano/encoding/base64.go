package encoding

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
)

func EncodeBase64(raw []byte) []byte {
	buff := new(bytes.Buffer)

	encoder := base64.NewEncoder(base64.StdEncoding, buff)
	defer encoder.Close()

	encoder.Write(raw)
	return buff.Bytes()
}

func DecodeBase64(raw []byte) ([]byte, error) {
	decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewBuffer(raw))

	b, err := ioutil.ReadAll(decoder)
	return b, err
}
