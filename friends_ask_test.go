package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAskFriends(t *testing.T) {
	is := require.New(t)

	setupTestServer()

	client := &http.Client{}
	resp := &AskFriendsResponse{}

	t.Run("Normal player asks friends - target id is self id", func(t *testing.T) {
		req := AskFriendsRequest{
			TargetID: testNormalProfile.ID,
			Message:  "hi, I want to be my friend",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, AskFriendsURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusBadRequest, code)
		is.False(resp.OK)
		is.Equal("cannot ask friends to self", resp.Error)
	})

	t.Run("Normal player asks friends - target id does not exist", func(t *testing.T) {
		req := AskFriendsRequest{
			TargetID: "not_exist_id",
			Message:  "hi, I want to be your friend",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, AskFriendsURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusInternalServerError, code)
		is.False(resp.OK)
		is.Equal("self profile or target profile does not exist", resp.Error)
	})

	t.Run("Normal player asks friends - request body is not valid", func(t *testing.T) {
		req := AskFriendsRequest{
			Message: "must not works",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, AskFriendsURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusBadRequest, code)
		is.False(resp.OK)
		is.Equal("target_id is required", resp.Error)
	})
}
