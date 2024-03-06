package music

type Scale struct {
	Name     string
	RootNote Note
	Notes    []Note
}

func NewScale(name string, rootNote Note, notes []Note) *Scale {
	return &Scale{Name: name, RootNote: rootNote, Notes: notes}
}

func NewMajorScale(rootNote Note) *Scale {
	return NewScale("Major", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MajorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MajorSixth(), rootNote.MajorSeventh()})
}

func NewMinorScale(rootNote Note) *Scale {
	return NewScale("Minor", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MinorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MinorSixth(), rootNote.MinorSeventh()})
}

func NewPentatonicScale(rootNote Note) *Scale {
	return NewScale("Pentatonic", rootNote, []Note{rootNote, rootNote.MinorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MinorSeventh()})
}

func NewHarmonicMinorScale(rootNote Note) *Scale {
	return NewScale("Harmonic Minor", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MinorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MinorSixth(), rootNote.MajorSeventh()})
}

func NewMelodicMinorScale(rootNote Note) *Scale {
	return NewScale("Melodic Minor", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MinorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MajorSixth(), rootNote.MajorSeventh()})
}

func NewDorianScale(rootNote Note) *Scale {
	return NewScale("Dorian", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MinorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MajorSixth(), rootNote.MinorSeventh()})
}

func NewPhrygianScale(rootNote Note) *Scale {
	return NewScale("Phrygian", rootNote, []Note{rootNote, rootNote.Sharp(), rootNote.MinorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MinorSixth(), rootNote.MinorSeventh()})
}

func NewLydianScale(rootNote Note) *Scale {
	return NewScale("Lydian", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MajorThird(), rootNote.Tritone(), rootNote.Fifth(), rootNote.MajorSixth(), rootNote.MajorSeventh()})
}

func NewMixolydianScale(rootNote Note) *Scale {
	return NewScale("Mixolydian", rootNote, []Note{rootNote, rootNote.Second(), rootNote.MajorThird(), rootNote.Fourth(), rootNote.Fifth(), rootNote.MajorSixth(), rootNote.MinorSeventh()})
}

func NewLocrianScale(rootNote Note) *Scale {
	return NewScale("Locrian", rootNote, []Note{rootNote, rootNote.Sharp(), rootNote.MinorThird(), rootNote.Fourth(), rootNote.Tritone(), rootNote.MinorSixth(), rootNote.MinorSeventh()})
}
