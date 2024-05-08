package github.com/efexplose/bits

import (
	"testing"
	"fmt"
	"log"
	"io"
	"os"
)

// To test Integer Encoding
func TestInt(t *testing.T) {
    bits := make([]bool, 64)
	n := 5
	size := 3
	decludeCount := 0

	encodeInt(&bits, size, decludeCount, n)
	i := decodeInt(&bits, 0, 3)
	j := decodeInt(&bits, 61, 64)

	fmt.Printf("Int:\n%s\n= %d = %d\n\n", Readable(bits), i, j)
}

func encodeInt(bits *[]bool, size int,decludeCount int, n int) {
	_, err := EncodeIntL(bits, size, decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
	_, err = EncodeIntR(bits, size, decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
}

func decodeInt(bits *[]bool, start int, end int) int {
	n, err := DecodeInt(bits, start, end)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// To test String Encoding
func TestString(t *testing.T) {
	s := "Hello, World!"
	bits := make([]bool, BitsNeededString(s)*2)
	decludeCount := 0
	
	encodeString(&bits, decludeCount, s)
	
	readable := ReadableGroup(bits, 8)

	difference(readHelloWorld(), readable)
}

func encodeString(bits *[]bool, decludeCount int, s string) {
	_, err := EncodeStringL(bits, decludeCount, s)
	if err != nil {
		log.Fatal(err)
	}
	_, err = EncodeStringR(bits, decludeCount, s)
	if err != nil {
		log.Fatal(err)
	}
}

func difference(reference, current string) {
	// colorize unmatching characters red
	var out string
	for i := 0; i < len(current); i++ {
		if reference[i] != current[i] {
			out += "\033[31m" + string(current[i]) + "\033[0m"
		} else {
			out += "\033[32m" + string(current[i]) + "\033[0m"	
		}
	}

	fmt.Printf("Current:\n%s\n\n", out)
}

func readHelloWorld() string {
	file, err := os.Open("./test_resources/hello_world.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	
	return string(b)
}

// To test Formatter
func TestFormatter(t *testing.T) {
	bits := make([]bool, 124)
	
	f := NewFormatter(&bits)
	f.Int(5, 0, 0)
	f.String("Hello, World!", 0)
	f.Bool(true, 0)
	f.Int(5, 3, 0)
	
	fmt.Printf("Formatter:\n%s\n\n", Readable(bits))
}

func TestEncodeAndDecodeBack(t *testing.T) {
	var bits []bool
	Encode(&bits, []byte("Hello, World!"))
	var bytes []byte
	Decode(&bytes, bits)
	fmt.Printf("Encode and Decode Back:\n%s\n\n", string(bytes))
}