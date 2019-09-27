package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerAskFriends create a friends action.
func HandlerAskFriends(rout *Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *AskFriendsResponse

		// Fetch session and valide it.
		sessID := r.URL.Query().Get("session")

		session, err := rout.DB.GetSession(sessID)
		if err != nil {
			resp = NewAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = session.Validate(); err != nil {
			resp = NewAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		// Parse request from body.
		defer r.Body.Close()
		var req AskFriendsRequest
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			resp = NewAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = req.Validate(); err != nil {
			resp = NewAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		if session.ProfileID == req.TargetID {
			resp = NewAskFriendsResponseErr(fmt.Errorf("cannot ask friends to self"))
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		// Fetch self profile and target profile from db.
		profiles, err := rout.DB.GetProfiles([]string{
			session.ProfileID,
			req.TargetID})
		if err != nil {
			resp = NewAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		// Check if target profile existed.
		if len(profiles) != 2 {
			resp = NewAskFriendsResponseErr(fmt.Errorf("self profile or target profile does not exist"))
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}
		selfProfile := profiles[0]
		targetProfile := profiles[1]

		// Check if the same action already existed.
		askFriends := NewAskFriendsByAsker(selfProfile.ID,
			targetProfile.ID,
			req.Message)
		existed, err := rout.DB.AskFriendsAlreadyExisted(askFriends)
		if err != nil {
			resp = NewAskFriendsResponseErr(fmt.Errorf("failed to check if ask friends action existed"))
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}
		if existed {
			resp = NewAskFriendsResponseErr(fmt.Errorf("ask friends action already existed"))
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		// Insert action in db.
		err = rout.DB.InserAskFriends(askFriends)
		if err != nil {
			resp = NewAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		// build response.
		resp = NewAskFriendsResponseOK(askFriends)
		writeResponse(w, resp, http.StatusOK)
	})
}
