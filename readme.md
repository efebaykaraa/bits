# Bits In Golang #

## Getting Started ##

### Installation ###
- Type `go get github.com/efexplose/bits` in the terminal
- Import the package as b `b import "github.com/efexplose/bits"`

### Basics ###
- As we don't have a `bit`/`bits` type in Golang, we can use `[]bool` to represent **bits**. False stands for 0 and true stands for 1. For example, line below represents **5** in binary: `101`
    ```go
    var bits = []bool{true, false, true}
    ```

### Bits & Bytes ###
- We can easily **convert** `[]bool` to `[]byte` in order to do **native tasks** such as write, send, etc.
    ```go
    var bits = []bool{true, false, true}
    var bytes []byte
    b.Decode(&bytes, bits)
    // bytes == []byte{5} = true
    ```
- We can also do the **opposite**, **convert** `[]byte` to `[]bool`.
    ```go
    var bytes = []byte{5}
    var bits []bool
    b.Encode(&bits, bytes)
    // bits == []bool{true, false, true} = true
    ```

### Coverting Other Types ###
- We can **convert** `int` to `[]bool`.
    ```go
    bits := make([]bool, 6) // Buffer to encode
    number := 5 // Number to encode
    size := BitsNeededInt(number) // Size limit, automaticly calculated
                                  // Same function will be also called if set to 0
    decludeCount := 0 // Bits to skip
    b.EncodeIntL(&bits, size, decludeCount, number) // Encode integer and alin to left
    b.EncodeIntR(&bits, size, decludeCount, number) // Encode integer and alin to right
    // bits == []bool{true, false, true, true, false, true} = true
    ```
- We can **convert** `string` to `[]bool`.
    ```go
    size := 8 * len(str) // Size limit, 8 bits for each character
    bits := size * 2 // Buffer to encode
    str := "Hello, World!" // String to encode
    decludeCount := 0 // Bits to skip
    b.EncodeStringL(&bits, size, decludeCount, str) // Encode string and alin to left
    b.EncodeStringR(&bits, size, decludeCount, str) // Encode string and alin to right
    // Too long to write down but basically it will encode hello world twice in the array
    ```

#### Note: *To be honest, I got tired of writing this markdown. So I must recommend you to check the binaries_test.go file.*

## Other Nessecary Stuff ##
- Also `string` and `int` can be converted to `[]bool` by using `DecodeString()` and `DecodeInt()` functions. **Examples included** in the test file.**
- Here is a formatter structure that works similar to `fmt.Sprintf` but easier to use. **Examples included** in the test file.
- Bitwise operations are possible by converting `[]bool` to `int`. **Examples not included** in the test file.

## Conclusion ##
- That's all for now. I hope I won't need to update the package anymore xD. I will porbably implement a simpler way to deal with bitwise operations but if you do it before I do, please pull request. You can also fork the project and take the responsibility off my shoulders.