package music_test

import (
	"reflect"
	"testing"

	"fretz/music"
)

func TestScales(t *testing.T) {
	// Define test cases for each scale
	tests := []struct {
		scaleFunc     func(music.Note) *music.Scale // Function to create scale
		rootNote      music.Note                    // Root note for the scale
		expectedNotes []music.Note                  // Expected notes for the scale
	}{
		{music.NewMajorScale, music.CNote, []music.Note{music.CNote, music.DNote, music.ENote, music.FNote, music.GNote, music.ANote, music.BNote}},
		{music.NewMinorScale, music.CNote, []music.Note{music.CNote, music.DNote, music.DSharpNote, music.FNote, music.GNote, music.GSharpNote, music.ASharpNote}},
		{music.NewPentatonicScale, music.ANote, []music.Note{music.ANote, music.CNote, music.DNote, music.ENote, music.GNote}},
		{music.NewHarmonicMinorScale, music.CNote, []music.Note{music.CNote, music.DNote, music.DSharpNote, music.FNote, music.GNote, music.GSharpNote, music.BNote}},
		{music.NewMelodicMinorScale, music.DNote, []music.Note{music.DNote, music.ENote, music.FNote, music.GNote, music.ANote, music.BNote, music.CSharpNote}},
		{music.NewDorianScale, music.ENote, []music.Note{music.ENote, music.FSharpNote, music.GNote, music.ANote, music.BNote, music.CSharpNote, music.DNote}},
		{music.NewPhrygianScale, music.FNote, []music.Note{music.FNote, music.FSharpNote, music.GSharpNote, music.ASharpNote, music.CNote, music.CSharpNote, music.DSharpNote}},
		{music.NewLydianScale, music.GNote, []music.Note{music.GNote, music.ANote, music.BNote, music.CSharpNote, music.DNote, music.ENote, music.FSharpNote}},
		{music.NewMixolydianScale, music.ANote, []music.Note{music.ANote, music.BNote, music.CSharpNote, music.DNote, music.ENote, music.FSharpNote, music.GNote}},
		{music.NewLocrianScale, music.BNote, []music.Note{music.BNote, music.CNote, music.DNote, music.ENote, music.FNote, music.GNote, music.ANote}},
	}

	// Iterate over test cases
	for _, tc := range tests {
		// Create scale using the scale function and root note
		scale := tc.scaleFunc(tc.rootNote)

		// Check scale name
		if scale.Name == "" {
			t.Errorf("Expected non-empty scale name, got empty name for %s scale with root note %s", scale.Name, tc.rootNote)
		}

		// Check scale root note
		if scale.RootNote != tc.rootNote {
			t.Errorf("Expected root note %s, got %s for %s scale", tc.rootNote, scale.RootNote, scale.Name)
		}

		// Check scale notes
		if !reflect.DeepEqual(scale.Notes, tc.expectedNotes) {
			t.Errorf("Expected notes %v, got %v for %s scale with root note %s", tc.expectedNotes, scale.Notes, scale.Name, tc.rootNote)
		}
	}
}
