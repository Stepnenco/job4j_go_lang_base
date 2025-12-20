package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {

	rsl := Add(1, 2)
	expected := 3

	assert.Equal(t, rsl, expected)

}
