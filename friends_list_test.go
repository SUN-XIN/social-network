package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListFriends(t *testing.T) {
	setupTestServer()

	is := require.New(t)

	client := &http.Client{}

	t.Run("whole work flow for friends relations for accept case", func(t *testing.T) {
		// ask friends.
		resp := &AskFriendsResponse{}
		req := AskFriendsRequest{
			TargetID: testSpecialProfile.ID,
			Message:  "hi, I want to be your friend",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, AskFriendsURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(resp.OK)
		is.Equal(req.Message, resp.Message)
		is.NotEmpty(resp.Created)

		// re-ask the same request.
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, AskFriendsURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusBadRequest, code)
		is.False(resp.OK)
		is.Equal("ask friends action already existed", resp.Error)

		// get ask friends request.
		respGet := &ListAskFriendsResponse{}
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, ListAskFriendsURL, testSessionSpecialProfile.ID),
			&req, respGet)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(respGet.OK)
		is.Len(respGet.Requests, 1)

		// accept ask friends request.
		respTreat := &Response{}
		reqTreat := TreatAskFriendsRequest{
			ActionID: respGet.Requests[0].ID,
			Action:   ActionAccept,
		}
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, TreatAskFriendsURL, testSessionSpecialProfile.ID),
			&reqTreat, respTreat)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(respTreat.OK)

		// list normal player's friends by normal player
		respList1 := &ListFriendsResponse{}
		reqList1 := ListFriendsRequest{
			TargetID: testNormalProfile.ID,
		}
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, ListFriendsURL, testSessionNormalProfile.ID),
			&reqList1, respList1)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(respList1.OK)
		is.Len(respList1.Friends, 1)
		is.Equal(respList1.Friends[0].FriendID, testSpecialProfile.ID)

		// list special player's friends by special player
		respList2 := &ListFriendsResponse{}
		reqList2 := ListFriendsRequest{
			TargetID: testSpecialProfile.ID,
		}
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, ListFriendsURL, testSessionSpecialProfile.ID),
			&reqList2, respList2)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(respList2.OK)
		is.Len(respList2.Friends, 1)
		is.Equal(respList2.Friends[0].FriendID, testNormalProfile.ID)

		// list normal player's friends by special player
		respList3 := &ListFriendsResponse{}
		reqList3 := ListFriendsRequest{
			TargetID: testNormalProfile.ID,
		}
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, ListFriendsURL, testSessionSpecialProfile.ID),
			&reqList3, respList3)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(respList3.OK)
		is.Len(respList3.Friends, 1)
		is.Equal(respList3.Friends[0].FriendID, testSpecialProfile.ID)

		// list special player's friends by normal player
		respList4 := &ListFriendsResponse{}
		reqList4 := ListFriendsRequest{
			TargetID: testSpecialProfile.ID,
		}
		code, err = sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, ListFriendsURL, testSessionNormalProfile.ID),
			&reqList4, respList4)
		is.NoError(err)
		is.Equal(http.StatusForbidden, code)
		is.False(respList4.OK)
		is.Equal("only Special player can fetch other's profile", respList4.Error)
	})

	// TODO: refuse case
}
