package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToByte(t *testing.T) {
	type example struct {
		Name string
	}
	text := "{\"Name\":\"test\"}"
	bytes := []byte(text)

	assert.EqualValues(t, bytes, ToByte(example{Name: "test"}))
}

func Test_Dump(t *testing.T) {
	type example struct {
		Name string
	}
	text := "{\"Name\":\"test\"}"

	assert.EqualValues(t, text, Dump(example{Name: "test"}))
}
