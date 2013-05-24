package bits

import (
    "testing"
)

type bytesPair struct {
    first []byte
    second []byte
    distance int
}

var distance_inputs []bytesPair = []bytesPair{
    bytesPair{
        []byte{ 0, 0, 0, 0, 0, 0 },
        []byte{ 1, 1, 1, 1, 1, 1 },
        6 },
    bytesPair{
        []byte{ 0, 0, 0, 0, 0, 0 },
        []byte{ 255, 255, 255, 255, 255, 255 },
        48 },
    bytesPair{
        []byte{ 255, 255, 255, 255, 255, 254 },
        []byte{ 255, 255, 255, 255, 255, 255 },
        1 },
    bytesPair{
        []byte{ 255, 255, 255, 255, 255, 127 },
        []byte{ 255, 255, 255, 255, 255, 255 },
        1 },
    }

// Number of bits in first that are not the same in second
func TestDistance(t *testing.T) {
    for _, pair := range distance_inputs {
        distance, err := Distance(pair.first, pair.second)
        if err != nil {
            t.Error("Distance returned error:", err)
        }

        if distance != pair.distance {
            t.Error("Distance(", pair.first, ",", pair.second, ") =",  distance, "!=", distance)
        }
    }
}
