package main

import (
	"regexp"
	"strings"
)

const (
	attributePrefix    = `1::Attributes`
	statusIDUnitPrefix = `1.7.4.2.`
	evDataPrefix       = `evData`
)

var (
	loginRegex = regexp.MustCompile(`^1::Login`)
	unitRegex  = regexp.MustCompile(`^1.7.4.2.\d+::evData`)
)

func IsAuthorization(m string) bool {
	r := regexp.MustCompile(`1::Authorize`)
	return r.Match([]byte(m))
}

func IsLogin(m string) bool {
	return loginRegex.Match([]byte(m))
}

func IsAttributes(m string) bool {
	return strings.HasPrefix(m, attributePrefix)
}

func IsUnitMessage(m string) bool {
	return unitRegex.Match([]byte(m))
}

func IsStatusIDUnitMessage(m string) bool {
	if !strings.HasPrefix(m, statusIDUnitPrefix) {
		return false
	}
	s := strings.TrimPrefix(m, statusIDUnitPrefix)
	r := strings.Split(s, "::")
	if len(r) < 2 {
		return false
	}
	return strings.HasPrefix(r[1], evDataPrefix)
}
