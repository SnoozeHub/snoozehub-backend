package main

import "testing"

func Test(t *testing.T) {
	t.Log(GenRandomString(100))
	t.Fail()
}
