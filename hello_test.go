package hello

import "testing"

func TestSayHello(t *testing.T) {
	if SayHello() != "Hello" {
		t.Fatal("SayHello must return Hello")
	}
}
