package flagparser

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// parse parses the given string
func (flag *Flag) parse(in string) (out interface{}, err error) {
	switch flag.Type {
	case "string":
		out = in
	case "int":
		out, err = strconv.ParseInt(in, 0, 0)
	case "duration":
		out, err = time.ParseDuration(in)
	}
	return out, err
}

// Parse parses the given slice of strings. Output is always given in the same order as NewFlagParser, followed by the remaining (non-captured) arguments as a slice of strings.
func (parser *FlagParser) Parse(in []string) (map[string]interface{}, error) {
	var itemsToRemove []int
	var err error
	var out = make(map[string]interface{})

	for _, arg := range parser.Flags {
		switch arg.Type {
		case "int":
			out[arg.Flags[0]] = 0
		case "string":
			out[arg.Flags[0]] = ""
		case "duration":
			duration, _ := time.ParseDuration("876600h")
			out[arg.Flags[0]] = duration
		case "bool":
			out[arg.Flags[0]] = false
		}
	}

	for i, arg := range in {
		if strings.HasPrefix(arg, "-") {
			arg = strings.TrimPrefix(arg, "-")
			for _, flag := range parser.Flags {
				var matched bool
				for _, flagName := range flag.Flags {
					if arg == flagName {
						matched = true
					}
				}
				if matched {
					if flag.Type == "bool" {
						out[flag.Flags[0]] = true
					} else {
						option, err := flag.parse(in[i+1])
						if err != nil {
							return out, err
						}
						switch option.(type) {
						case int, int64:
							out[flag.Flags[0]] = option.(int64)
						case time.Duration:
							out[flag.Flags[0]] = option.(time.Duration)
						default:
							out[flag.Flags[0]] = option.(string)
						}
						itemsToRemove = append(itemsToRemove, i, i+1)
					}
				}
			}
		}
	}
	// clean up the output
	sort.Ints(itemsToRemove)
	reverseAny(itemsToRemove)
	for _, i := range itemsToRemove {
		in = removeFromSlice(in, i)
	}
	out["rest"] = in
	return out, err
}

func removeFromSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
