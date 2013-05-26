package repeatingxor

import (
    "testing"
)

type repeating_xor_input struct {
    plaintext []byte
    key []byte
    ciphertext []byte
}

var repeating_xor_inputs []repeating_xor_input = []repeating_xor_input{
    repeating_xor_input{
        plaintext:  []byte{ 127, 127, 128, 127, 127, 128},
        key:        []byte{ 128, 128, 127 },
        ciphertext: []byte{ 255, 255, 255, 255, 255, 255},
    },
    repeating_xor_input{
        plaintext:  []byte{ 0, 0, 0, 0, 0, 0, 0 },
        key:        []byte{ 1, 2, 3, 4 },
        ciphertext: []byte{ 1, 2, 3, 4, 1, 2, 3 },
    },
}

var repeating_xor_errors []repeating_xor_input = []repeating_xor_input{
    repeating_xor_input{
        plaintext:  []byte{ },
        key:        []byte{ 0 },
    },
    repeating_xor_input{
        plaintext:  []byte{ 0 },
        key:        []byte{ },
    },
}


func TestRepeatingXor(t *testing.T) {
    for _, i := range repeating_xor_inputs {
        if ciphertext, err := RepeatingXor(i.plaintext, i.key); err != nil {
            t.Error("RepeatingXor(", i.plaintext, i.key, ") returned error:", err)

        } else if string(ciphertext) != string(i.ciphertext) {
            t.Error(ciphertext, "!=", i.ciphertext)
        }

        if plaintext, err := RepeatingXor(i.ciphertext, i.key); err != nil {
            t.Error("RepeatingXor(", i.ciphertext, i.key, ") returned error:", err)

        } else if string(plaintext) != string(i.plaintext) {
            t.Error(plaintext, "!=", i.plaintext)
        }
    }

    for _, i := range repeating_xor_errors {
        if _, err := RepeatingXor(i.plaintext, i.key); err == nil {
            t.Error("RepeatingXor(", i.plaintext, i.key, ") did not return error")
        }
    }
}
