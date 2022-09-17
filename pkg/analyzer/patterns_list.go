package analyzer

import (
	"regexp"
)

type PatternsList []*regexp.Regexp

// newPatternsList parses slice of strings to a slice of compiled regular
// expressions.
func newPatternsList(in []string) (PatternsList, error) {
	list := make(PatternsList, len(in))

	for i, str := range in {
		re, err := strToRegexp(str)
		if err != nil {
			return nil, err
		}

		list[i] = re
	}

	return list, nil
}

// MatchesAny matches provided string against all regexps in a slice.
func (l PatternsList) MatchesAny(str string) bool {
	for _, r := range l {
		if r.MatchString(str) {
			return true
		}
	}

	return false
}
