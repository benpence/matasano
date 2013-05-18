package bits

func RepeatingXor(plaintext []byte, key []byte) []byte {
    if len(plaintext) == 0 || len(key) == 0 {
        plaintext_ := make([]byte, len(plaintext))
        for i, b := range plaintext {
            plaintext_[i] = b
        }
        return plaintext_
    }

    ciphertext := make([]byte, len(plaintext))

    // TODO: Use existing Xor function instead?
    for i, b := range plaintext {
        ciphertext[i] = b ^ key[i % len(key)]
    }

    return ciphertext
}
