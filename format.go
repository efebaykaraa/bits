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

func (f *Formatter) Int(n int, decludeCount int) {
	f.Position += decludeCount
	err := EncodeIntL(f.Buffer, f.Position + decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
}

func (f *Formatter) String(s string, decludeCount int) {
	f.Position += decludeCount
	err := EncodeStringL(f.Buffer, f.Position + decludeCount, s)
	if err != nil {
		log.Fatal(err)
	}
}

func (f *Formatter) Bool(b bool, decludeCount int) {
	f.Position += decludeCount
	(*f.Buffer)[f.Position] = b
}