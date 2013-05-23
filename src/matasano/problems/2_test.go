package problems

import (
	"testing"

	"matasano/bits"
	"matasano/encoding"
)

const input2a string = "1c0111001f010100061a024b53535009181c"
const input2b string = "686974207468652062756c6c277320657965"
const output2 string = "746865206b696420646f6e277420706c6179"

func Test2(t *testing.T) {
    raw1, err := encoding.DecodeHex([]byte(input2a))
    if err != nil {
        t.Error("Decoding", input2a, "failed")
    }

    raw2, err := encoding.DecodeHex([]byte(input2b))
    if err != nil {
        t.Error("Decoding", input2b, "failed")
    }

    result, err := bits.Xor(raw1, raw2)
    if err != nil {

    }

    resultEncoded := encoding.EncodeHex(result)
    if string(resultEncoded) != output2 {
        t.Error(string(resultEncoded), "!=", output2)
	}
}
