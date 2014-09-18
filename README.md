Go implementation of the Matasano Crypto Challenges (http://cryptopals.com)

# Source

The source in `src/matasano` is organized as follows:

- `bits/` general bit-operation
- `encoding/` base64 and hex encoding and decoding
- `xorbyte/` encrypting and forcibly decrypting message blocks XOR'd by the same byte
- `repeatingxor/` encrypting and forcibly decrypting message blocks XOR'd by a bytestring
- `aes/` encrypting with AES and detecting AES
- `problems/` matasano-specific problem numbers

# Testing

Each problem solution is organized into a go `testing` function in `src/matasano/problems/`. To run all the tests, execute:

    $ GOPATH="`pwd`:$GOPATH" go test matasano/problems

To examine the tests for a particular problem, please read the respective `src/matasano/problems/NUMBER_test.go` file.
