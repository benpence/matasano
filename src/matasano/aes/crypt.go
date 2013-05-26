package aes

import (
    "crypto/cipher"
    "crypto/aes"
)

func Encrypt(text, key []byte) ([]byte, error) {
    // Create cipher and pad text
    cipher, padded, err := cipherAndPaddedText(text, key)
    if err != nil {
        return nil, err
    }

    keysize := len(key)

    // Encrypt each block
    for i := 0; i < len(padded) / keysize; i += keysize {
        cipher.Encrypt(padded[i: i + keysize], padded[i: i + keysize])
    }

    return padded[:len(text)], nil
}

func Decrypt(text, key []byte) ([]byte, error) {
    // Create cipher and pad text
    cipher, padded, err := cipherAndPaddedText(text, key)
    if err != nil {
        return nil, err
    }

    keysize := len(key)

    // Decrypt each block
    for i := 0; i < len(padded); i += keysize {
        cipher.Decrypt(padded[i: i + keysize], padded[i: i + keysize])
    }

    return padded[:len(text)], nil
}

func cipherAndPaddedText(text, key []byte) (cipher.Block, []byte, error) {
    // Create cipher
    cipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, nil, err
    }

    // Round ciphertext size up to nearest block
    var padded []byte
    if len(text) % 16 != 0 {
        padded = make([]byte, len(text) + 16 - (len(text) % 16))
    } else {
        padded = make([]byte, len(text))
    }

    // Copy text into padded container
    copy(padded, text)

    return cipher, padded, nil
}
