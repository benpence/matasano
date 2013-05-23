package matasano

import (
	"testing"

	"matasano/encoding"
)

const input1 string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const output1 string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func Test1(t *testing.T) {
    raw, err := encoding.DecodeHex([]byte(input1))
    if err != nil {
        t.Error("Decoding", input1, "failed")
    }

    if encoded := encoding.EncodeBase64(raw); string(encoded) != output1 {
        t.Error(string(encoded), "!=", output1)
	}
}
