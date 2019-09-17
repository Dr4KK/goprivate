package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	greets := HelloWorld()
	assert.Equal(t, Greeting, greets)
}
