package flagparser

import (
	"reflect"
	"testing"
)

func TestFlagParser(t *testing.T) {
	var (
		flagParser *FlagParser
		input      []string
		expected   = make(map[string]interface{})
		output     = make(map[string]interface{})
	)

	flagParser, _ = NewFlagParser(Int("min"), Int("max"), Duration("d"), String("x"), Bool("bool"))

	input = []string{"hello", "world", "-d", "10m", "-max", "5", "world2", "-min", "3", "-x", "hello", "sdfasdff", "-bool"}
	expected = map[string]interface{}{"rest": []string{"hello", "world", "sdfasdff"}, "bool": true, "x": "hello", "max": int64(5), "min": int64(3), "d": 600000000000}

	output, err := flagParser.Parse(input)
	if err != nil {
		t.Fatalf("Parse: %v = %v; got error %v", input, expected, err)
	}
	if reflect.DeepEqual(output, expected) {
		t.Errorf("Parse: %v = %v; got %v", input, expected, output)
	}
}

func TestFlagParser2(t *testing.T) {
	var (
		flagParser *FlagParser
		input      []string
		expected   = make(map[string]interface{})
		output     = make(map[string]interface{})
	)

	flagParser, _ = NewFlagParser(Int("min"), Int("max", "mx"), Duration("d"), String("x"), Bool("bool"))

	input = []string{"hello", "world", "-mx", "5", "world2", "-x", "hello", "sdfasdff"}
	expected = map[string]interface{}{"rest": []string{"hello", "world", "sdfasdff"}, "bool": true, "x": "hello", "max": int64(5), "min": int64(3), "d": 3155760000000000000}

	output, err := flagParser.Parse(input)
	if err != nil {
		t.Fatalf("Parse: %v = %v; got error %v", input, expected, err)
	}
	if reflect.DeepEqual(output, expected) {
		t.Errorf("Parse: %v = %v; got %v", input, expected, output)
	}
}
