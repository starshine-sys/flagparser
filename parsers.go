package flagparser

import (
	"fmt"
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
			out[arg.Flag] = 0
		case "string":
			out[arg.Flag] = ""
		case "duration":
			duration, _ := time.ParseDuration("876600h")
			out[arg.Flag] = duration
		}
	}

	for i, arg := range in {
		fmt.Println(i)
		if len(in) > i+1 && strings.HasPrefix(arg, "-") {
			arg = strings.TrimPrefix(arg, "-")
			for _, flag := range parser.Flags {
				if arg == flag.Flag {
					fmt.Println(arg, i, flag.Type)
					option, err := flag.parse(in[i+1])
					if err != nil {
						return out, err
					}
					fmt.Printf("%v => %T, %v\n", flag.Flag, option, option)
					switch option.(type) {
					case int, int64:
						out[flag.Flag] = option.(int64)
					case time.Duration:
						out[flag.Flag] = option.(time.Duration)
					case bool:
						out[flag.Flag] = option.(bool)
					default:
						out[flag.Flag] = option.(string)
					}
					itemsToRemove = append(itemsToRemove, i, i+1)
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
