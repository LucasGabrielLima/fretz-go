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
