package aes

import "errors"

// Find most likely AES 128 ECB encrypted ciphertext by summing the number of
//   repititions of 16 byte blocks and returning the ciphertext with the most
func FindEcb(ciphertexts [][]byte, blocksize int) ([]byte, error) {
    if len(ciphertexts) == 0 {
        return nil, errors.New("ciphertexts must not be empty")
    } else if blocksize <= 0 {
        return nil, errors.New("blocksize must be > 0")
    }

    // Defaults in case there are no ciphertexts with repititions
    mostCount := 0
    bestCiphertext := ciphertexts[0]

    for _, ciphertext := range ciphertexts {
        totalCount := 0

        // Sum up repititions. 
        for _, count := range groupBlocksByValue(ciphertext, blocksize) {
            if count > 1 {
                // Ciphertexts whose reptitions are concentrated on fewer
                //   blocks will score higher
                totalCount += count - 1
            }
        }

        // Update if more reptitions
        if totalCount > mostCount {
            mostCount = totalCount
            bestCiphertext = ciphertext
        }
    }

    return bestCiphertext, nil
}

// Group blocks of size blocksize together and count their occurences
func groupBlocksByValue(blocks []byte, blocksize int) map[string] int {
    // TODO: Research overhead/pitfalls of string conversion
    groups := make(map[string] int)

    // NOTE: Does not count leftover bytes smaller than a block
    for i := 0; len(blocks) - i >= blocksize; i += blocksize {
        block := string(blocks[i: i + blocksize])

        // Already seen this block?
        if _, ok := groups[block]; ok {
            groups[block]++

        // New block
        } else {
            groups[block] = 1
        }
    }

    return groups
}
