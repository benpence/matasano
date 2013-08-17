package repeatingxor

import (
	"testing"
)

type findKeysizeInput struct {
    plaintext string
    key []byte
}

const plaintext string =
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here" +
    "This is plaintext written for use in the go test that I will be performing in order to test this function that we have here"

var findKeysizeInputs []findKeysizeInput = []findKeysizeInput{
    findKeysizeInput{
        plaintext,
        []byte{ 1, 2, 3, 4 },
    },
    findKeysizeInput{
        plaintext,
        []byte{ 1, 2, 3, 4, 5, 6, 7, 8 },
    },
    findKeysizeInput{
        plaintext,
        []byte{ 1, 171, 18, 126, 91, 138, 200, 219 },
    },
}

func TestFindKeysize(t *testing.T) {
    for _, i := range findKeysizeInputs {

        ciphertext, err := Crypt([]byte(i.plaintext), i.key)
        if err != nil {
            t.Error("Encrypting", i.plaintext, "failed")
        }

        keysize, err := FindKeysize(ciphertext, 2, 40, 5)
        if err != nil {
            t.Error("repeatingxor.Crypt returned error:", err)
        }

        if keysize != len(i.key) {
            t.Error(keysize, "!=", len(i.key))
        }
    }
}
