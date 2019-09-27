package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		title    string
		input    []string
		expected []string
	}{
		{
			title:    "no any duplicate",
			input:    []string{"1", "2", "3"},
			expected: []string{"1", "2", "3"},
		},
		{
			title:    "empty input",
			input:    []string{},
			expected: []string{},
		},
		{
			title:    "2 duplicates",
			input:    []string{"1", "2", "1"},
			expected: []string{"1", "2"},
		},
		{
			title:    "3 duplicates",
			input:    []string{"1", "1", "1"},
			expected: []string{"1"},
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()

			got := RemoveDuplicate(testCase.input)
			is := require.New(t)
			is.Equal(testCase.expected, got)
		})
	}
}
