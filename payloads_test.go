package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListFriendsRequestValidate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		title       string
		input       ListFriendsRequest
		expected    ListFriendsRequest
		expectedErr error
	}{
		{
			title: "good case",
			input: ListFriendsRequest{
				TargetID: "id1",
				Pagination: Pagination{
					Limit: 5,
				},
			},
			expected: ListFriendsRequest{
				TargetID: "id1",
				Pagination: Pagination{
					Limit: 5,
				},
			},
			expectedErr: nil,
		},
		{
			title:       "target_id is missing",
			input:       ListFriendsRequest{},
			expected:    ListFriendsRequest{},
			expectedErr: fmt.Errorf("target_id is required"),
		},
		{
			title: "limit value is missing",
			input: ListFriendsRequest{
				TargetID: "id2",
			},
			expected: ListFriendsRequest{
				TargetID: "id2",
				Pagination: Pagination{
					Limit: DEFAULT_PAGE_LIMIT,
				},
			},
			expectedErr: nil,
		},
		{
			title: "limit value is too big",
			input: ListFriendsRequest{
				TargetID: "id3",
				Pagination: Pagination{
					Limit: MAX_PAGE_LIMIT + 1,
				},
			},
			expected: ListFriendsRequest{
				TargetID: "id3",
				Pagination: Pagination{
					Limit: MAX_PAGE_LIMIT,
				},
			},
			expectedErr: nil,
		},
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.title, func(t *testing.T) {
			t.Parallel()

			err := testCase.input.Validate()
			is := require.New(t)
			is.Equal(testCase.expectedErr, err)
			is.Equal(testCase.expected, testCase.input)
		})
	}
}
