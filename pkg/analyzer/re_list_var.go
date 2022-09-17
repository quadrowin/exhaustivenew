package analyzer

import (
	"errors"
	"fmt"
	"regexp"
)

// Структура используется для чтения аргументов командной строки
type reListVar struct {
	values *PatternsList
}

func (v *reListVar) Set(value string) error {
	re, err := strToRegexp(value)
	if err != nil {
		return err
	}

	*v.values = append(*v.values, re)

	return nil
}

func (v *reListVar) String() string {
	return ""
}

var ErrEmptyPattern = errors.New("pattern can't be empty")

func strToRegexp(str string) (*regexp.Regexp, error) {
	if str == "" {
		return nil, ErrEmptyPattern
	}

	re, err := regexp.Compile(str)
	if err != nil {
		return nil, fmt.Errorf("unable to compile %s as regular expression: %w", str, err)
	}

	return re, nil
}
