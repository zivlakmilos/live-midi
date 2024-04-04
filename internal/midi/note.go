package midi

import "math"

type NoteType int

const (
	NoteTypeNormal NoteType = iota
	NoteTypeRised
	NoteTypeLowered
)

var normalizeMapping = [12]uint8{
	0,
	0,
	1,
	1,
	2,
	3,
	3,
	4,
	4,
	5,
	5,
	6,
}

func Normalize(note uint8) (uint8, int8) {
	octave := int8(math.Floor((float64(note) - 60.0) / 12.0))
	note = note % 12

	return normalizeMapping[note], octave
}
