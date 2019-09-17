package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	greets := HelloWorld()
	assert.Equal(t, Greeting, greets)
}
