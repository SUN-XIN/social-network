package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ListFriendsURL     = "/test/list_friends"
	GetProfilesURL     = "/test/get_profile"
	AskFriendsURL      = "/test/ask_friends"
	ListAskFriendsURL  = "/test/list_ask_friends"
	TreatAskFriendsURL = "/test/treat_ask_friends"
	CreateReportURL    = "/test/create_report"

	PrefixURL = "http://localhost:8081"
)

var (
	testServerIsOn = false
)

func setupTestServer() {
	if testServerIsOn {
		return
	}

	rout := NewRouter()
	err := seedGetProfiles(rout)
	if err != nil {
		panic(err)
	}

	http.Handle(ListAskFriendsURL, HandlerListAskFriends(rout))
	http.Handle(AskFriendsURL, HandlerAskFriends(rout))
	http.Handle(TreatAskFriendsURL, HandlerTreatAskFriends(rout))
	http.Handle(ListFriendsURL, HandlerListFriends(rout))
	http.Handle(GetProfilesURL, HandlerGetProfiles(rout))
	http.Handle(CreateReportURL, HandlerCreateReport(rout))

	go func() {
		err := http.ListenAndServe(":8081", nil)
		panic(err)
	}()

	testServerIsOn = true
}

func sendPostRequest(client *http.Client,
	url string,
	objReq, objResp interface{}) (int, error) {
	b, err := json.Marshal(objReq)
	if err != nil {
		return -1, fmt.Errorf("Failed Marshal request body: %+v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return -1, fmt.Errorf("Failed create request: %+v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("Failed send request: %+v", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(objResp); err != nil {
		return -1, fmt.Errorf("Failed decode response: %+v", err)
	}

	return resp.StatusCode, nil
}
