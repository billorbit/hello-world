package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	expected := "hello world"
	actural := hello()
	assert.Equal(t, expected, actural)
}
