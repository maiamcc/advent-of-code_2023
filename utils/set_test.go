package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	is := IntSet{
		1: struct{}{},
		4: struct{}{},
	}

	expected := IntSet{
		1: struct{}{},
		4: struct{}{},
		6: struct{}{},
	}

	is.Add(6)
	assert.Equal(t, expected, is)

	// Adding an existing element should be a no-op
	is.Add(4)
	assert.Equal(t, expected, is)
}

func TestNewIntSet(t *testing.T) {
	expected := IntSet{
		1: struct{}{},
		4: struct{}{},
		6: struct{}{},
	}
	assert.Equal(t, expected, NewIntSet([]int{6, 4, 1}))
}

func TestIntSet_Contains(t *testing.T) {
	is := IntSet{
		1: struct{}{},
		4: struct{}{},
	}
	assert.True(t, is.Contains(4))
	assert.False(t, is.Contains(100))
}

func TestIntSet_Rm(t *testing.T) {
	is := IntSet{
		1: struct{}{},
		4: struct{}{},
		6: struct{}{},
	}
	expected := IntSet{
		1: struct{}{},
		4: struct{}{},
	}

	is.Rm(6)
	assert.Equal(t, expected, is)

	// Removing an element not present in the set is a no-op
	is.Rm(600)
	assert.Equal(t, expected, is)
}
