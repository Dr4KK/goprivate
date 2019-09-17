package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	greets := HelloWorld()
	assert.Equal(t, Greeting, greets)
}
