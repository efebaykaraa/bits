package bits

import (
	"testing"
	"fmt"
	"log"
	"io"
	"os"
	"github.com/efexplose/bits/needed"
	"github.com/efexplose/bits/format"
	"github.com/efexplose/bits/encode"
	"github.com/efexplose/bits/decode"
)

const hw = "Hello, World!"

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
	_, err := encode.IntL(bits, size, decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
	_, err = encode.IntR(bits, size, decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
}

func decodeInt(bits *[]bool, start int, end int) int {
	n, err := decode.Int(bits, start, end)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// To test String Encoding
func TestString(t *testing.T) {
	bits := make([]bool, needed.String(hw)*2)
	decludeCount := 0
	
	encode.StringL(&bits, decludeCount, hw)
	encode.StringR(&bits, decludeCount, hw)
	
	readable := ReadableGroup(bits, 8)

	difference(readHelloWorld(), readable)
}

func EncodeString(bits *[]bool, decludeCount int, s string) {
	_, err := encode.StringL(bits, decludeCount, s)
	if err != nil {
		log.Fatal(err)
	}
	_, err = encode.StringR(bits, decludeCount, s)
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
func TestFormatters(t *testing.T) {
	bits := make([]bool, 128)
	
	f := format.NewFormatter(&bits)
	f.Int(5, 0, 0)
	f.String("Hello, World!", 0)
	f.Bool(true, 0)
	f.Int(5, 3, 0)
	
	fmt.Printf("Formatter:\n%s\n\n", Readable(bits))

	d := format.NewDeformatterL(&bits)
	i := d.Int(3, 0)
	s := d.String(needed.String("Hello, World!"), 0)
	b := d.Bool(0)
	j := d.Int(3, 0)
	
	fmt.Printf("Deformatter:\n%d\n%s\n%t\n%d\n\n", i, s, b, j)
}

func TestEncodeAndDecodeBack(t *testing.T) {
	var bits []bool
	Encode(&bits, []byte("Hello, World!"))
	var bytes []byte
	Decode(&bytes, bits)
	fmt.Printf("Encode and Decode Back:\n%s\n\n", string(bytes))
}