package repeatingxor

import (
    "fmt"
    "errors"

    "matasano/bits"
    "matasano/xorbyte"
)

const SAMPLE_BLOCKS int = 5
const KEYSIZE_LOW int = 2
const KEYSIZE_HIGH int = 40

func FindKey(ciphertext []byte) ([]byte, error) {
    keysize, err := FindKeysize(ciphertext, KEYSIZE_LOW, KEYSIZE_HIGH, SAMPLE_BLOCKS)
    if err != nil {
        return nil, err
    }

    fmt.Println(keysize)

    // TODO: How does padding affect single xor cracking?
    transposed, err := bits.Transpose(ciphertext, keysize)
    if err != nil {
        return nil, err
    }

    blocksTotal := len(transposed) / keysize
    key := make([]byte, blocksTotal)

    for i := 0; i < blocksTotal; i++ {
        block := transposed[i * keysize: (i + 1) * keysize]
        key[i] = xorbyte.FindByte(block, xorbyte.EnglishScorer)
    }

    return key, nil
}

// Determine likely blocksize by finding the blocksize that minimizes the 
//   amount of differing bits between blocks in a sample of the ciphertext
// NOTE: assumes len(ciphertext) >= highest * sampleBlocks
func FindKeysize(ciphertext []byte, lowest, highest, sampleBlocks int) (int, error) {
    if highest < lowest || lowest < 1 {
        return -1, errors.New("highest > lowest > 0 must hold")

    } else if sampleBlocks * highest > len(ciphertext) {
        return -1, errors.New("len(ciphertext) > sampleBlocks * highest must hold")
    }

    var totalDistances float64 = float64((sampleBlocks * (sampleBlocks - 1)) / 2)
    if totalDistances == 0 {
        totalDistances = 1
    }

    var least_distance float64 = -1

    // Default value
    best_keysize := lowest

    for keysize := lowest; keysize <= highest; keysize++ {
        // Average all distances between first 'sampleBlocks' blocks in ciphertext
        sample := ciphertext[: sampleBlocks * keysize]

        var average_distance float64
        for blockPair := range pairs(sample, keysize) {
            distance, _ := bits.Distance(blockPair.first, blockPair.second)
            average_distance += float64(distance) / float64(keysize) / totalDistances
        }
        /*

        distance , err := bits.Distance(ciphertext[:keysize], ciphertext[keysize: 2 * keysize])
        if err != nil {
            panic(err)
        }
        average_distance := float64(distance) / float64(keysize)
        */

        fmt.Println(keysize, average_distance)

        // Found better keysize?
        if least_distance == -1 || average_distance < least_distance {
            least_distance = average_distance
            best_keysize = keysize
        }

    }

    return best_keysize, nil
}

type pair struct {
    first []byte
    second []byte
}

// A channel of all possible block pairs in text
func pairs(text []byte, blocksize int) chan pair {
    ch := make(chan pair)

    go func() {
        // TODO: Last few leftover bytes?
        for i := 0; len(text) - i >= blocksize; i += blocksize {
            for j := i + blocksize; len(text) - j >= blocksize; j += blocksize {
                ch <- pair{ text[i: i + blocksize], text[j: j + blocksize] }
            }
        }

        close(ch)
    }()

    return ch
}


