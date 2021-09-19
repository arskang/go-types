package gotypes

import (
	"fmt"
	"strings"
)

type String string

func (s *String) Len() int {
	if s != nil {
		return len(string(*s))
	}
	return 0
}

func (s *String) ToLower() string {
	if s != nil {
		return strings.ToLower(string(*s))
	}
	return ""
}

func (s *String) ToUpper() string {
	if s != nil {
		return strings.ToUpper(string(*s))
	}
	return ""
}

func (s *String) ToPrint(message string) {
	if s != nil {
		fmt.Println(message, *s)
	}
}
