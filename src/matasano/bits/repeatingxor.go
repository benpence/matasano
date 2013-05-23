package bits

import (
    "errors"
)

// Xor text with key repeating to fit the length of text
func RepeatingXor(text []byte, key []byte) ([]byte, error) {
    // Paramters empty?
    if len(text) == 0 || len(key) == 0 {
        return nil, errors.New("text and key must not be empty")
    }

    ciphertext := make([]byte, len(text))

    for i, b := range text {
        ciphertext[i] = b ^ key[i % len(key)]
    }

    return ciphertext, nil
}
