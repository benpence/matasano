package matasano

import (
	"testing"

	"matasano/bits"
	"matasano/encoding"
)

const input1 string = "1c0111001f010100061a024b53535009181c"
const input2 string = "686974207468652062756c6c277320657965"
const output string = "746865206b696420646f6e277420706c6179"

func Test2(t *testing.T) {
    raw1, err := encoding.DecodeHex([]byte(input1))
    if err != nil {
        t.Error("Decoding", input1, "failed")
    }

    raw2, err := encoding.DecodeHex([]byte(input2))
    if err != nil {
        t.Error("Decoding", input2, "failed")
    }

    result, err := bits.Xor(raw1, raw2)
    if err != nil {

    }

    resultEncoded := encoding.EncodeHex(result)
    if string(resultEncoded) != output {
        t.Error(string(resultEncoded), "!=", output)
	}
}
