package main

import (
	"regexp"
	"strings"
)

const (
	statusIDUnitPrefix = `1.7.4.2.`
	evDataPrefix       = `evData`
	authorize          = `1::Authorize`
	loginPrefix        = `1::Login`
)

var (
	authorizationCompiled = regexp.MustCompile(`^1::Authorize`)
	loginRegex            = regexp.MustCompile(`^1::Login`)
	unitRegex             = regexp.MustCompile(`^1.7.4.2.\d+::evData`)
)

func IsAuthorization(m string) bool {
	r := regexp.MustCompile(`1::Authorize`)
	return r.Match([]byte(m))
}

func IsAuthorizationCompiled(m string) bool {
	return authorizationCompiled.Match([]byte(m))
}

func IsAuthorizationString(m string) bool {
	return strings.Contains(m, `1::Authorize`)
}

func IsAuthorizationStringConstant(m string) bool {
	return strings.Contains(m, authorize)
}

func IsLogin(m string) bool {
	return loginRegex.Match([]byte(m))
}

func IsLoginString(m string) bool {
	return strings.HasPrefix(m, loginPrefix)
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
