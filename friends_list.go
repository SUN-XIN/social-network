package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerListFriends returns the friends of a given profile by pagination.
func HandlerListFriends(rout *Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *ListFriendsResponse

		// Fetch session and valide it.
		sessID := r.URL.Query().Get("session")
		session, err := rout.DB.GetSession(sessID)
		if err != nil {
			resp = NewListFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = session.Validate(); err != nil {
			resp = NewListFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		// Parse request from body.
		defer r.Body.Close()
		var request ListFriendsRequest
		if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
			resp = NewListFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		if err = request.Validate(); err != nil {
			resp = NewListFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		// Fetch self profile from db.
		profiles, err := rout.DB.GetProfiles([]string{session.ProfileID})
		if err != nil {
			resp = NewListFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		// Check profile's status.
		selfProfile := profiles[0]
		if selfProfile.Status != StatusSpecial &&
			session.ProfileID != request.TargetID {
			resp = NewListFriendsResponseErr(fmt.Errorf("only Special player can fetch other's profile"))
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		// Fetch friends from db.
		friends, err := rout.DB.ListFriends(request.TargetID)
		if err != nil {
			resp = NewListFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		// build response.
		resp = NewListFriendsResponseOK(friends, Pagination{Limit: request.Limit})
		writeResponse(w, resp, http.StatusOK)
	})
}
