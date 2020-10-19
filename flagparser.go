package flagparser

// FlagParser parses the flags
type FlagParser struct {
	Flags []*Flag
}

// Flag is a single flag, meant to be used in FlagParser
type Flag struct {
	Flags []string
	Type  string
}

// NewFlagParser returns a new FlagParser instance
func NewFlagParser(flags ...*Flag) (parser *FlagParser, err error) {
	parser = &FlagParser{
		Flags: flags,
	}
	return parser, err
}

// Int returns an int parser
func Int(flags ...string) *Flag {
	return &Flag{
		Flags: flags,
		Type:  "int",
	}
}

// String returns a string parser
func String(flags ...string) *Flag {
	return &Flag{
		Flags: flags,
		Type:  "string",
	}
}

// Duration returns a duration parser
func Duration(flags ...string) *Flag {
	return &Flag{
		Flags: flags,
		Type:  "duration",
	}
}

// Bool returns a bool parser
func Bool(flags ...string) *Flag {
	return &Flag{
		Flags: flags,
		Type:  "bool",
	}
}
