package main

import (
	"net/http"
)

// HandlerListAskFriends fetch all my ask friends requests.
// TODO: can filter the response
// for example
// 1. only return the ask friends requests not yet treated
// 2. only return the ask friends requests not yet read
func HandlerListAskFriends(rout *Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *ListAskFriendsResponse

		// Fetch session and valide it.
		sessID := r.URL.Query().Get("session")
		session, err := rout.DB.GetSession(sessID)
		if err != nil {
			resp = NewListAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = session.Validate(); err != nil {
			resp = NewListAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		// Fetch my ask friends requests from db.
		askFriendsRequests, err := rout.DB.GetMyAskFriends(session.ProfileID)
		if err != nil {
			resp = NewListAskFriendsResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		// TODO: update all fetched ask friends request as read
		// AskFriends.ReadAt = now

		resp = NewListAskFriendsResponseOK(askFriendsRequests)
		writeResponse(w, resp, http.StatusOK)
	})
}
