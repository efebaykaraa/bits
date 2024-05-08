package bits

import (
	"log"
)

type Formatter struct {
	Buffer 		*[]bool
	Position 	int
}

func NewFormatter(bits *[]bool) *Formatter {
	return &Formatter{
		Buffer: bits,
		Position: 0,
	}
}

func (f *Formatter) Skip(n int) {
	f.Position += n
}

func (f *Formatter) Int(n int, size int, decludeCount int) {
	n, err := EncodeIntL(f.Buffer, size, f.Position + decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
	f.Position += decludeCount + n
}

func (f *Formatter) String(s string, decludeCount int) {
	n, err := EncodeStringL(f.Buffer, f.Position + decludeCount, s)
	if err != nil {
		log.Fatal(err)
	}
	f.Position += decludeCount + n
}

func (f *Formatter) Bool(b bool, decludeCount int) {
	(*f.Buffer)[f.Position] = b
	f.Position += decludeCount + 1
}