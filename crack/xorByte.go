package crack

import (
	"errors"
    "math"

	"matasano/bits"
)

// Calculate the distance by summing the difference between each relative frequency
//   and each corresponding desired frequencies
func distance(relativeFrequencies []float32, desiredFrequencies []float32) float64 {
    var d float64 = 0

    for i := 0; i < len(desiredFrequencies); i++ {
        d += math.Abs(float64(relativeFrequencies[i] - desiredFrequencies[i]))
    }

    return d
}

// Finds byte most likely to have been xor'd with every byte in ciphertext
// NOTE: Only operates on single bytes, not unicode characters
func XorByte(ciphertext []byte, desiredFrequencies []float32) (byte, error) {
	if len(desiredFrequencies) != 256 {
		return 0, errors.New("Desired frequencies must have length 256")
	}

	var (
		closestByte byte  = 0
		leastDistance float64 = -1
	)

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
		frequencyCount := make([]byte, 256) // Slice is zeroed
		for _, p := range plaintext {
			frequencyCount[p]++
		}

		// Calculate distance between measured and desired relative frequencies
        relativeFrequencies := make([]float32, 256)
        for byt := 0; byt < 256; byt++ {
            relativeFrequencies[byt] = float32(frequencyCount[byt]) / float32(len(ciphertext))
        }

        distance := distance(relativeFrequencies, desiredFrequencies)

		// Update the most likely character
		if distance < leastDistance || leastDistance == -1 {
			closestByte = byte(byteValue)
			leastDistance = distance
		}
	}

	return closestByte, nil
}
