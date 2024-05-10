package format

import (
	"log"
	"github.com/efexplose/bits/decode"
)

type Deformatter struct {
	Buffer 		*[]bool
	Position 	int
}

func NewDeformatterL(bits *[]bool) *Deformatter {
	return &Deformatter{
		Buffer: bits,
		Position: 0,
	}
}

func (f *Deformatter) Skip(n int) {
	f.Position += n
}

func (f *Deformatter) Int(size int, decludeCount int) int {
	v, err := decode.Int(f.Buffer, size, f.Position + decludeCount)
	if err != nil {
		log.Fatal(err)
	}
	f.Position += decludeCount + size
	return v
}

func (f *Deformatter) String(size int, decludeCount int) string {
	println(f.Position + decludeCount + size > len(*f.Buffer))
	s := decode.String(f.Buffer, size, f.Position + decludeCount)
	f.Position += decludeCount + size 
	return s
}

func (f *Deformatter) Bool(decludeCount int) bool {
	b := (*f.Buffer)[f.Position + decludeCount]
	f.Position += decludeCount + 1
	return b
}