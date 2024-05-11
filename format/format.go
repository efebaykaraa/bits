package format

import (
	"log"
	"github.com/efexplose/bits/encode"
	"github.com/efexplose/bits/needed"
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
	n, err := encode.IntL(f.Buffer, size, f.Position + decludeCount, n)
	if err != nil {
		log.Fatal(err)
	}
	f.Position += decludeCount + n
}

func (f *Formatter) IntR(n int, size int, decludeCount int) {
	neededInt := needed.Int(n)
	alignedOffset := size - neededInt + f.Position + decludeCount
	n, err := encode.IntL(f.Buffer, size, alignedOffset, neededInt)
	if err != nil {
		log.Fatal(err)
	}
	f.Position += decludeCount + n
}

func (f *Formatter) String(s string, decludeCount int) {
	n, err := encode.StringL(f.Buffer, f.Position + decludeCount, s)
	if err != nil {
		log.Fatal(err)
	}
	f.Position += decludeCount + n
}

func (f *Formatter) Bool(b bool, decludeCount int) {
	(*f.Buffer)[f.Position] = b
	f.Position += decludeCount + 1
}