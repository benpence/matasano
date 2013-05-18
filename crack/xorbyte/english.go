package xorbyte

import "math"

var EnglishScorer Scorer = Scorer(functionScorer(englishScore))

func englishScore(text []byte) float64 {
    total := 0
    counts := make([]int, 256)
    var score float64

    for _, c := range text {
        switch {
        // Lowercase characters and space
        case (97 <= c && c <= 122) || c == 32:
            counts[int(c)]++
            total++

        // Uppercase
        case 65 <= c && c <= 90:
            counts[int(c + 32)]++
            total++

        // Numbers, symbols, newline & carriage return
        case (33 <= c && c <= 64) || (91 <= c && 96 <= c) || (123 <= c && 126 <= c) || c == 10 || c == 13:
            continue

        // Invalid character -> invalid message
        default:
            // TODO: Infinity?
            return 1000000
        }
    }

    // Calculate the score by summing the difference between each measured frequency
    //   and each corresponding desired frequencies
    for c, count := range counts {
        if frequency, ok := englishFrequencies[string(c)]; ok {
            score += math.Abs(float64(frequency - float32(count) / float32(total)))
        }
    }

    return score
}

var englishFrequencies map[string]float32 = map[string]float32 {
    " ": 0.1918182, "a": 0.0651738, "b": 0.0124248,
    "c": 0.0217339, "d": 0.0349835, "e": 0.1041442,
    "f": 0.0197881, "g": 0.0158610, "h": 0.0492888,
    "i": 0.0558094, "j": 0.0009033, "k": 0.0050529,
    "l": 0.0331490, "m": 0.0202124, "n": 0.0564513,
    "o": 0.0596302, "p": 0.0137645, "q": 0.0008606,
    "r": 0.0497563, "s": 0.0515760, "t": 0.0729357,
    "u": 0.0225134, "v": 0.0082903, "w": 0.0171272,
    "x": 0.0013692, "y": 0.0145984, "z": 0.0007836,
}
