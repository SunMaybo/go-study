package test

import "testing"

type M struct {
	A *int
}

func TestEscape(t *testing.T) {
   var a *int
   m:=&M{}
   m.A=a
}
