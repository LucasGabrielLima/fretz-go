package music

type Neck struct {
	Strings []GuitarString
}

func NewNeck() *Neck {
	strings := make([]GuitarString, 6)
	stringsRoots := [...]Note{ENote, BNote, GNote, DNote, ANote, ENote}
	for i, root := range stringsRoots {
		strings[i] = *NewGuitarString(i, root, 24)
	}
	return &Neck{Strings: strings}
}

func (n Neck) Format() {
	for _, guitarString := range n.Strings {
		guitarString.Format()
	}
}
