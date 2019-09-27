package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateReport(t *testing.T) {
	is := require.New(t)

	setupTestServer()

	client := &http.Client{}
	resp := &Response{}

	t.Run("good case", func(t *testing.T) {
		req := CreateReportRequest{
			TargetID: testSpecialProfile.ID,
			Motive:   "test",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, CreateReportURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusOK, code)
		is.True(resp.OK)
	})

	t.Run("session is not found", func(t *testing.T) {
		req := CreateReportRequest{
			TargetID: testSpecialProfile.ID,
			Motive:   "test",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, CreateReportURL, "bad_session"),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusInternalServerError, code)
		is.False(resp.OK)
	})

	t.Run("bad request", func(t *testing.T) {
		req := CreateReportRequest{
			TargetID: testSpecialProfile.ID,
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, CreateReportURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusBadRequest, code)
		is.False(resp.OK)
		is.Equal("motive is required", resp.Error)
	})

	t.Run("target profile is not found", func(t *testing.T) {
		req := CreateReportRequest{
			TargetID: "not_exist",
			Motive:   "test",
		}
		code, err := sendPostRequest(client,
			fmt.Sprintf("%s%s?session=%s", PrefixURL, CreateReportURL, testSessionNormalProfile.ID),
			&req, resp)
		is.NoError(err)
		is.Equal(http.StatusNotFound, code)
		is.False(resp.OK)
		is.Equal("target profile is not found in db", resp.Error)
	})
}
