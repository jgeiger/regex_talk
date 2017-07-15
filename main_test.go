package main

import "testing"

func BenchmarkIsAuthorizationRegexp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorizationRegexp("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsAuthorizationRegexpCompiled(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorizationRegexpCompiled("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsAuthorizationRegexpCompiledMatchString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorizationRegexpCompiledMatchString("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsAuthorizationStringsContains(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorizationStringsContains("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsAuthorizationStringsContainsConstant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorizationStringsContainsConstant("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsAuthorizationStringsPrefix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorizationStringsPrefix("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsUnitMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsUnitMessage("1.7.4.2.12345::evData")
	}
}

func BenchmarkIsStatusIDUnitMessage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsStatusIDUnitMessage("1.7.4.2.12345::evData")
	}
}
