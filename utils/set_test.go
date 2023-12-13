package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIntSet(t *testing.T) {
	expected := IntSet{
		1: struct{}{},
		4: struct{}{},
		6: struct{}{},
	}
	assert.Equal(t, expected, NewIntSet([]int{6, 4, 1}))
}

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

func TestNewGenericSet(t *testing.T) {
	expected := Set[int]{
		1: struct{}{},
		4: struct{}{},
		6: struct{}{},
	}
	assert.Equal(t, expected, NewSet(6, 4, 1))
}

func TestGenericSet_Add(t *testing.T) {
	set := NewSet(1, 4)

	expected := Set[int]{
		1: struct{}{},
		4: struct{}{},
		6: struct{}{},
	}

	set.Add(6)
	assert.Equal(t, expected, set)

	// Adding an existing element should be a no-op
	set.Add(4)
	assert.Equal(t, expected, set)
}

func TestGenericSet_Contains(t *testing.T) {
	set := NewSet(1, 4)
	assert.True(t, set.Contains(4))
	assert.False(t, set.Contains(100))
}

func TestGenericSet_Rm(t *testing.T) {
	set := NewSet(1, 4)
	expected := Set[int]{
		1: struct{}{},
		4: struct{}{},
	}

	set.Rm(6)
	assert.Equal(t, expected, set)

	// Removing an element not present in the set is a no-op
	set.Rm(600)
	assert.Equal(t, expected, set)
}
