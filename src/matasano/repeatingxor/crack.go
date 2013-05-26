package repeatingxor

import (
    "matasano/bits"
    "matasano/xorbyte"
)

func FindKey(ciphertext []byte) []byte {
    keysize := findKeysize(ciphertext, 2, 40, 4)

    transposed := transpose(ciphertext, keysize)
    blocksTotal := len(transposed) / keysize
    key := make([]byte, blocksTotal)

    for i := 0; i < blocksTotal; i++ {
        block := transposed[i * keysize: (i + 1) * keysize]
        key[i] = xorbyte.FindByte(block, xorbyte.EnglishScorer)
    }

    return key
}

func findKeysize(ciphertext []byte, lowest, highest, sampleBlocks int) int {
    var comparisons float64 = float64((sampleBlocks * (sampleBlocks + 1)) / 2)
    var least_distance float64 = -1
    best_keysize := lowest

    for keysize := lowest; keysize <= highest; keysize++ {
        // Average all distances between first 'sampleBlocks'
        var average_distance float64 = 0
        sample := ciphertext[: sampleBlocks * keysize]

        for pair := range pairs(sample, keysize) {
            distance, _ := bits.Distance(pair.first, pair.second)
            average_distance += (float64(distance) / float64(keysize)) / comparisons
        }

        // Found better keysize?
        if least_distance == -1 || average_distance < least_distance {
            least_distance = average_distance
            best_keysize = keysize
        }
    }

    return best_keysize
}

type pair struct {
    first []byte
    second []byte
}

func pairs(text []byte, keysize int) chan pair {
    ch := make(chan pair)

    go func() {
        // TODO: Last few leftover bytes?
        for i := 0; len(text) - i >= keysize; i += keysize {
            for j := i + keysize; len(text) - j >= keysize; j += keysize {
                ch <- pair{ text[i: i + keysize], text[j: j + keysize] }
            }
        }

        close(ch)
    }()

    return ch
}

func transpose(ciphertext []byte, keysize int) []byte {
    // Ensure transposition will be square matrix
    // TODO: How does this affect single xor cracking?
    length := len(ciphertext)
    if length % keysize != 0 {
        length += keysize - (length % keysize)
    }

    transposed := make([]byte, length)

    for i, b := range ciphertext {
        transposed[((i % keysize) * keysize) + i / keysize] = b
    }

    return transposed
}
