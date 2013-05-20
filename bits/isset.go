package bits

func IsSet(b byte, i uint) bool {
    i = 7 - (i % 8)
    return ((b & (1 << i)) >> i) == 1
}
