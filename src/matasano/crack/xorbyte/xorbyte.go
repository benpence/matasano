package xorbyte

import "matasano/bits"

type byteAndScore struct {
    byte
    Score float64
}

// Finds ciphertext most likely to have been xor'd with a key of all one byte
// NOTE: Only operates on single bytes, not unicode characters
func FindCiphertext(ciphertexts [][]byte, scorer Scorer) ([]byte, byte) {
    var (
		closestCiphertext []byte = make([]byte, 256)
        closestByte byte = 0
		lowestScore float64 = -1
    )

    for _, ciphertext := range ciphertexts {
        byteValue, score := xorByte(ciphertext, scorer)

        // Update the most likely ciphertext
        if score < lowestScore || lowestScore == -1 {
            closestCiphertext = ciphertext
            closestByte = byteValue
            lowestScore = score
        }
    }

    return closestCiphertext, closestByte
}

// Finds byte most likely to have been xor'd with every byte in ciphertext
// NOTE: Only operates on single bytes, not unicode characters
func FindByte(ciphertext []byte, scorer Scorer) byte {
    closestByte, _ := xorByte(ciphertext, scorer)
    return closestByte
}

func xorByte(ciphertext []byte, scorer Scorer) (byte, float64) {
	var (
		closestByte byte = 0
		lowestScore float64 = -1
	)

    for byteScore := range bytesAndScores(ciphertext, scorer) {
		// Update the most likely character
		if byteScore.Score < lowestScore || lowestScore == -1 {
			closestByte = byteScore.byte
			lowestScore = byteScore.Score
		}
    }

    return closestByte, lowestScore
}

func bytesAndScores(ciphertext []byte, scorer Scorer) chan byteAndScore {
    ch := make(chan byteAndScore)

    go func() {
        key := make([]byte, len(ciphertext))

        // Test each possible value of byte
        for byteValue := 0; byteValue < 256; byteValue++ {

            // Fill key with byteValue
            for i := 0; i < len(key); i++ {
                key[i] = byte(byteValue)
            }

            // Xor ciphertext with key (p == (p xor k) xor k)
            plaintext, _ := bits.Xor(ciphertext, key)

            // Frequency count
            ch <- byteAndScore{
                byte(byteValue),
                scorer.Score(plaintext),
            }
        }

        close(ch)
	}()

	return ch
}
