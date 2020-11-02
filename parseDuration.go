package flagparser

// Code taken from YAGPDB.xyz, https://github.com/jonas747/yagpdb
// Code licensed under the MIT license
// Source: https://github.com/jonas747/yagpdb/blob/ad6ea1f14ab463bdf3c5e29c6e9c2e12cfce675d/commands/util.go

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"emperror.dev/errors"
)

// Parses a time string like 1day3h
func parseDuration(str string) (time.Duration, error) {
	var dur time.Duration

	currentNumBuf := ""
	currentModifierBuf := ""

	// Parse the time
	for _, v := range str {
		// Ignore whitespace
		if unicode.Is(unicode.White_Space, v) {
			continue
		}

		if unicode.IsNumber(v) {
			// If we reached a number and the modifier was also set, parse the last duration component before starting a new one
			if currentModifierBuf != "" {
				if currentNumBuf == "" {
					currentNumBuf = "1"
				}
				d, err := parseDurationComponent(currentNumBuf, currentModifierBuf)
				if err != nil {
					return d, err
				}

				dur += d

				currentNumBuf = ""
				currentModifierBuf = ""
			}

			currentNumBuf += string(v)

		} else {
			currentModifierBuf += string(v)
		}
	}

	if currentNumBuf != "" {
		d, err := parseDurationComponent(currentNumBuf, currentModifierBuf)
		if err != nil {
			return dur, errors.WrapIf(err, "not a duration")
		}

		dur += d
	}

	return dur, nil
}

func parseDurationComponent(numStr, modifierStr string) (time.Duration, error) {
	parsedNum, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0, err
	}

	parsedDur := time.Duration(parsedNum)

	if strings.HasPrefix(modifierStr, "s") {
		parsedDur = parsedDur * time.Second
	} else if modifierStr == "" || (strings.HasPrefix(modifierStr, "m") && (len(modifierStr) < 2 || modifierStr[1] != 'o')) {
		parsedDur = parsedDur * time.Minute
	} else if strings.HasPrefix(modifierStr, "h") {
		parsedDur = parsedDur * time.Hour
	} else if strings.HasPrefix(modifierStr, "d") {
		parsedDur = parsedDur * time.Hour * 24
	} else if strings.HasPrefix(modifierStr, "w") {
		parsedDur = parsedDur * time.Hour * 24 * 7
	} else if strings.HasPrefix(modifierStr, "mo") {
		parsedDur = parsedDur * time.Hour * 24 * 30
	} else if strings.HasPrefix(modifierStr, "y") {
		parsedDur = parsedDur * time.Hour * 24 * 365
	} else {
		return parsedDur, errors.New("couldn't figure out what '" + numStr + modifierStr + "` was")
	}

	return parsedDur, nil

}
