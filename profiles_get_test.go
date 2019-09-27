package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProfiles(t *testing.T) {
	is := require.New(t)

	setupTestServer()

	client := &http.Client{}
	resp := &GetProfilesResponse{}

	t.Run("Normal player gets his profile", func(t *testing.T) {
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, GetProfilesURL, testSessionNormalProfile.ID),
			[]string{testNormalProfile.ID}, resp)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(resp.OK)
		is.Len(resp.Profiles, 1)
		is.Equal(testNormalProfile.ID, resp.Profiles[0].ID)
		is.Equal("getprofiletest_name1", resp.Profiles[0].Name)
	})

	t.Run("Normal player gets other's profile.", func(t *testing.T) {
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, GetProfilesURL, testSessionNormalProfile.ID),
			[]string{testSpecialProfile.ID}, resp)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(resp.OK)
		is.Len(resp.Profiles, 1)
		is.Equal(testSpecialProfile.ID, resp.Profiles[0].ID)
		is.Equal("getprofiletest_name2", resp.Profiles[0].Name)
	})

	t.Run("One profile is not found.", func(t *testing.T) {
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, GetProfilesURL, testSessionNormalProfile.ID),
			[]string{testSpecialProfile.ID, "not_exist_id"}, resp)
		is.NoError(err)
		is.Equal(http.StatusPartialContent, code)
		is.False(resp.OK)
		is.Equal("only 1 profiles are fetched for 2 ids", resp.Error)
	})

	t.Run("Request body is empty.", func(t *testing.T) {
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, GetProfilesURL, testSessionNormalProfile.ID),
			[]string{}, resp)
		is.NoError(err)
		is.Equal(http.StatusBadRequest, code)
		is.False(resp.OK)
		is.Equal("target profile ids are required", resp.Error)
	})
}
