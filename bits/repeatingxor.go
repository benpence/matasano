package bits

func RepeatingXor(text []byte, key []byte) []byte {
    if len(text) == 0 || len(key) == 0 {
        text_ := make([]byte, len(text))
        for i, b := range text {
            text_[i] = b
        }
        return text_
    }

    ciphertext := make([]byte, len(text))

    // TODO: Use existing Xor function instead?
    for i, b := range text {
        ciphertext[i] = b ^ key[i % len(key)]
    }

    return ciphertext
}
