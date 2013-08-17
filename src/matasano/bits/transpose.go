package bits

import (
    "errors"
)

// Transposes text as a two dimensional matrix with rows of size rowLength
// NOTE: Length of transposed matrix will be rounded up to nearest multiple of
//   rowLength
func Transpose(text []byte, rowLength int) ([]byte, error) {
    if rowLength < 1 {
        return nil, errors.New("rowLength must be greater than 0")
    }

    // Ensure transposition will be rectangular matrix
    length := len(text)
    if length % rowLength != 0 {
        length += rowLength - (length % rowLength)
    }

    transposed := make([]byte, length)

    columnLength := len(transposed) / rowLength
    for i, b := range text {
        newRow := i % rowLength
        newColumn := i / rowLength

        transposed[newRow * columnLength + newColumn] = b
    }

    return transposed, nil
}
