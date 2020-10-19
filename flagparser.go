package flagparser

// FlagParser parses the flags
type FlagParser struct {
	Flags []*Flag
}

// Flag is a single flag, meant to be used in FlagParser
type Flag struct {
	Flag string
	Type string
}

// NewFlagParser returns a new FlagParser instance
func NewFlagParser(flags ...*Flag) (parser *FlagParser, err error) {
	parser = &FlagParser{
		Flags: flags,
	}
	return parser, err
}

// Int returns an int parser
func Int(flag string) *Flag {
	return &Flag{
		Flag: flag,
		Type: "int",
	}
}

// String returns a string parser
func String(flag string) *Flag {
	return &Flag{
		Flag: flag,
		Type: "string",
	}
}

// Duration returns a duration parser
func Duration(flag string) *Flag {
	return &Flag{
		Flag: flag,
		Type: "duration",
	}
}

// Bool returns a bool parser
func Bool(flag string) *Flag {
	return &Flag{
		Flag: flag,
		Type: "bool",
	}
}
