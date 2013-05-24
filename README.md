# Source

The source in `src/matasano` is organized as follows:

- `bits/` for encryption, decryption, and general bit-operation
- `encoding/` for base64 and hex encoding and decoding
- `crack/` for encryption detection and 'forced' ciphertext decryption
- `problems/` for test code

# Testing

Each problem solution is organized into a go `testing` function in `src/matasano/problems/`. To run all the tests, execute:

    $ GOPATH="`pwd`:$GOPATH" go test matasano/problems

To examine the tests for a particular problem, please read the respective `src/matasano/problems/NUMBER_test.go` file.
