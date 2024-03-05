package music

import "fmt"

type GuitarString struct {
	stringNumber int
	rootNote     Note
	frets        []Fret
}

func NewGuitarString(stringNumber int, rootNote Note, numPositions int) *GuitarString {
	frets := make([]Fret, numPositions)
	for i := 0; i < numPositions; i++ {
		noteNumber := (int(rootNote) + i) % 12
		fret := NewFret(Note(noteNumber), stringNumber, i)
		frets[i] = *fret
	}

	return &GuitarString{stringNumber: stringNumber, rootNote: rootNote, frets: frets}
}

func (g GuitarString) Format() {
	fmt.Printf("%s: ", g.rootNote.String())
	for _, fret := range g.frets {
		fret.Format()
	}

	fmt.Print("\n")
}
