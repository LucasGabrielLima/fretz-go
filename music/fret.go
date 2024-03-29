package music

import "fmt"

type Fret struct {
	note         Note
	stringNumber int
	position     int
}

func NewFret(note Note, stringNumber int, position int) *Fret {
	return &Fret{note: note, stringNumber: stringNumber, position: position}
}

func (f Fret) Format() {
	fmt.Printf(" - |")
}

func (f Fret) FormatWithNote() {
	fmt.Printf(" %v |", f.note.String())
}

func (f Fret) FormatWithPosition() {
	fmt.Printf(" %d |", f.position)
}
