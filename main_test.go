package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

// simple unit test
func TestSelectionSort(t *testing.T) {
	expectedResult := []int{1, 2, 3, 4, 5, 6}

	result := selectionSort([]int{5, 3, 1, 2, 4, 6})

	assert.Equal(t, expectedResult, result, "Should be equal")
}

// benchmark tests
func BenchmarkTestSelectionSort(b *testing.B) {
	n := make([]int, 5)

	// randomly generate numbers and add to the list
	for i := 0; i < 5; i++ {
		n[i] = rand.Intn(100)
	}

	for i := 0; i < b.N; i++ {
		selectionSort(n)
	}
}

// fuzz tests
func FuzzTestSelectionSort(f *testing.F) {
	n := make([]int, 5)

	// randomly generate numbers and add to the list
	for i := 0; i < 5; i++ {
		n[i] = rand.Intn(100)
	}

	f.Add(-3)
	f.Fuzz(func(t *testing.T, a int) {
		selectionSort(n)
	})
}

// Table driven tests to help reduce repetition. Keep tests DRY.
func TestSelectionSortTableDriven(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input []int
		want  []int
	}{
		// the table itself
		{"use standard set of unsorted numbers", []int{34, 45, 9}, []int{9, 34, 45}},
		{"use a sorted set as input", []int{1, 2, 3}, []int{1, 2, 3}},
		{"only one input", []int{1}, []int{1}},
		{"empty input", []int{}, []int{}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := selectionSort(tt.input)
			assert.Equal(t, tt.want, ans, "Should be equal")
		})
	}
}

// By default, tests are run sequentially. The method Parallel() indicates that all tests should be run in parallel.
// The GOMAXPROCS environment defines how many tests can run in parallel at one time.
// By default, this number is equal to the number of CPUs.
func TestSelectionSortParallel(t *testing.T) {
	t.Run("Test in Parallel - Random unsorted numbers", func(t *testing.T) {
		t.Parallel()
		result := selectionSort([]int{3, 5, 6})
		expected := []int{3, 5, 6}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
		}
	})
	t.Run("Test in Parallel - Random unsorted negative numbers", func(t *testing.T) {
		t.Parallel()
		result := selectionSort([]int{-3, -5, -6})
		expected := []int{-6, -5, -3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
		}
	})
}
