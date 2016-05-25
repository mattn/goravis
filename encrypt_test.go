package main

import (
	"os"
	"testing"
)

func TestEncrypt(t *testing.T) {
	content := os.Getenv("foo")
	if content != "bar" {
		t.Fatalf("should be %v but %v:", "bar", content)
	}
	content = os.Getenv("bar")
	if content != "baz" {
		t.Fatalf("should be %v but %v:", "baz", content)
	}
}
