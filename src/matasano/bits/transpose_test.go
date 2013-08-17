package bits

import (
    "testing"
)

type transposeInput struct {
    input []byte
    rowLength int
    output []byte
}

var transposeInputs []transposeInput = []transposeInput {
    transposeInput{
        []byte{ 1, 2, 3 },
        3,
        []byte{ 1,
                2,
                3 },
    },
    transposeInput{
        []byte{ 1, 2,
                3     },
        2,
        []byte{ 1, 3,
                2, 0 },
    },
    transposeInput{
        []byte{ 1, 2, 3, 4, 5, 6 },
        8,
        []byte{ 1,
                2,
                3,
                4,
                5,
                6,
                0,
                0 },
    },
    transposeInput{
        []byte{ 1, 2, 3,
                4, 5, 6,
                7, 8, 9 },
        3,
        []byte{ 1, 4, 7,
                2, 5, 8,
                3, 6, 9 },
    },
}

func TestTranspose(t *testing.T) {
    for _, i := range transposeInputs {
        if output, err := Transpose(i.input, i.rowLength); err != nil {
            t.Error("Transpose(", i.input, ",", i.rowLength, ") returned err:", err)

        } else if string(output) != string(i.output) {
            t.Error(output, "!=", i.output)
        }
    }
}
