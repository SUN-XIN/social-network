package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetID(t *testing.T) {
	f := Friends{
		MyID:     "id1",
		FriendID: "id2",
	}

	f.SetID()

	is := require.New(t)
	is.Equal("id1_id2", f.ID)
}
