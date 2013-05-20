package bits

import (
    "errors"
)

func Distance(first, second []byte) (int, error) {
    if len(first) != len(second) {
        return 0, errors.New("Byte slices must be same length")
    }

    count := 0
    for i, b := range first {
        for j := 0; j < 8; j++ {
            if IsSet(b, uint(j)) != IsSet(second[i], uint(j)) {
                count++
            }
        }
    }

    return count, nil
}
