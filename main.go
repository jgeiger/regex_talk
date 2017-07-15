package main

import (
	"regexp"
	"strings"
)

const (
	statusIDUnitPrefix = `1.7.4.2.`
	evDataPrefix       = `evData`
	authorize          = `1::Authorize`
)

var (
	authorizationCompiled = regexp.MustCompile(`^1::Authorize`)
	unitRegex             = regexp.MustCompile(`^1.7.4.2.\d+::evData`)
)

func IsAuthorizationRegexp(m string) bool {
	r := regexp.MustCompile(`1::Authorize`)
	return r.Match([]byte(m))
}

func IsAuthorizationRegexpCompiled(m string) bool {
	return authorizationCompiled.Match([]byte(m))
}

func IsAuthorizationRegexpCompiledMatchString(m string) bool {
	return authorizationCompiled.MatchString(m)
}

func IsAuthorizationStringsContains(m string) bool {
	return strings.Contains(m, `1::Authorize`)
}

func IsAuthorizationStringsContainsConstant(m string) bool {
	return strings.Contains(m, authorize)
}

func IsAuthorizationStringsPrefix(m string) bool {
	return strings.HasPrefix(m, `1::Authorize`)
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
