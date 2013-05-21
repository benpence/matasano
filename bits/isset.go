package bits

// Returns whether the ith bit of b (big endian) set
func IsSet(b byte, i uint) bool {
    i = 7 - (i % 8)
    return ((b & (1 << i)) >> i) == 1
}
