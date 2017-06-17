package main

import "testing"

func BenchmarkIsAuthorization(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAuthorization("1::Authorize(12345, ABCD)")
	}
}

func BenchmarkIsLogin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsLogin("1::Login(password)")
	}
}

func BenchmarkIsAttributes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsAttributes("1::Attributes(fan)")
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
