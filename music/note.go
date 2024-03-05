package music

type Note int

const (
	ANote Note = iota
	ASharpNote
	BNote
	CNote
	CSharpNote
	DNote
	DSharpNote
	ENote
	FNote
	FSharpNote
	GNote
	GSharpNote
)

func (n Note) String() string {
	return [...]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}[n]
}

func (n Note) ShiftSemitones(semitones int) Note {
	return Note(int(n)+semitones) % 12
}

func (n Note) Second() Note {
	return n.ShiftSemitones(2)
}

func (n Note) MinorThird() Note {
	return n.ShiftSemitones(3)
}

func (n Note) MajorThird() Note {
	return n.ShiftSemitones(4)
}

func (n Note) Fourth() Note {
	return n.ShiftSemitones(5)
}

func (n Note) Tritone() Note {
	return n.ShiftSemitones(6)
}

func (n Note) Fifth() Note {
	return n.ShiftSemitones(7)
}

func (n Note) MinorSixth() Note {
	return n.ShiftSemitones(8)
}

func (n Note) MajorSixth() Note {
	return n.ShiftSemitones(9)
}

func (n Note) MinorSeventh() Note {
	return n.ShiftSemitones(10)
}

func (n Note) MajorSeventh() Note {
	return n.ShiftSemitones(11)
}

func (n Note) Octave() Note {
	return n.ShiftSemitones(12)
}

func (n Note) MinorChord() [3]Note {
	return [3]Note{n, n.MinorThird(), n.Fifth()}
}

func (n Note) MajorChord() [3]Note {
	return [3]Note{n, n.MajorThird(), n.Fifth()}
}

func (n Note) DiminishedChord() [3]Note {
	return [3]Note{n, n.MinorThird(), n.Tritone()}
}

func (n Note) AugmentedChord() [3]Note {
	return [3]Note{n, n.MajorThird(), n.MinorSixth()}
}

func (n Note) SeventhChord() [4]Note {
	return [4]Note{n, n.MajorThird(), n.Fifth(), n.MinorSeventh()}
}

func (n Note) MajorSeventhChord() [4]Note {
	return [4]Note{n, n.MajorThird(), n.Fifth(), n.MajorSeventh()}
}

func (n Note) MinorSeventhChord() [4]Note {
	return [4]Note{n, n.MinorThird(), n.Fifth(), n.MinorSeventh()}
}

func (n Note) MinorMajorSeventhChord() [4]Note {
	return [4]Note{n, n.MinorThird(), n.Fifth(), n.MajorSeventh()}
}

func (n Note) DiminishedSeventhChord() [4]Note {
	return [4]Note{n, n.MinorThird(), n.Tritone(), n.MinorSeventh()}
}

func (n Note) HalfDiminishedSeventhChord() [4]Note {
	return [4]Note{n, n.MinorThird(), n.Tritone(), n.MinorSeventh()}
}

func (n Note) AugmentedSeventhChord() [4]Note {
	return [4]Note{n, n.MajorThird(), n.MinorSixth(), n.MinorSeventh()}
}

func (n Note) AugmentedMajorSeventhChord() [4]Note {
	return [4]Note{n, n.MajorThird(), n.MinorSixth(), n.MajorSeventh()}
}

func (n Note) NinthChord() [5]Note {
	return [5]Note{n, n.MajorThird(), n.Fifth(), n.MinorSeventh(), n.Second()}
}

func (n Note) MajorNinthChord() [5]Note {
	return [5]Note{n, n.MajorThird(), n.Fifth(), n.MajorSeventh(), n.Second()}
}

func (n Note) MinorNinthChord() [5]Note {
	return [5]Note{n, n.MinorThird(), n.Fifth(), n.MinorSeventh(), n.Second()}
}

func (n Note) MinorMajorNinthChord() [5]Note {
	return [5]Note{n, n.MinorThird(), n.Fifth(), n.MajorSeventh(), n.Second()}
}

func (n Note) DiminishedNinthChord() [5]Note {
	return [5]Note{n, n.MinorThird(), n.Tritone(), n.MinorSeventh(), n.Second()}
}

func (n Note) HalfDiminishedNinthChord() [5]Note {
	return [5]Note{n, n.MinorThird(), n.Tritone(), n.MinorSeventh(), n.Second()}
}

func (n Note) AugmentedNinthChord() [5]Note {
	return [5]Note{n, n.MajorThird(), n.MinorSixth(), n.MinorSeventh(), n.Second()}
}

func (n Note) AugmentedMajorNinthChord() [5]Note {
	return [5]Note{n, n.MajorThird(), n.MinorSixth(), n.MajorSeventh(), n.Second()}
}
