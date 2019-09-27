package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerTreatAskFriends treats a ask friends action.
func HandlerTreatAskFriends(rout *Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp Response

		// Fetch session and valide it.
		sessID := r.URL.Query().Get("session")

		session, err := rout.DB.GetSession(sessID)
		if err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = session.Validate(); err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		// Parse request from body.
		defer r.Body.Close()
		var req TreatAskFriendsRequest
		if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = req.Validate(); err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		// Check if action exists in db.
		askFriends, err := rout.DB.GetAskFriends(req.ActionID)
		if err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if askFriends.Receiver != session.ProfileID {
			// Should never happen.
			resp = NewResponseErr(fmt.Errorf("you can only treat your ask friends request"))
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		switch req.Action {
		case ActionUnkown:
			// Should never happen.
			resp = NewResponseErr(fmt.Errorf("treat action is unknown"))
			writeResponse(w, resp, http.StatusBadRequest)
			return
		case ActionAccept:
			// insert a new line friends into db.
			friends := NewFriends(askFriends.Asker, askFriends.Receiver)
			err = rout.DB.InserFriends(friends)
			if err != nil {
				resp = NewResponseErr(err)
				writeResponse(w, resp, http.StatusInternalServerError)
				return
			}
			// notify the asker.
			askFriends.AcceptByReceiver()
			err = rout.DB.UpdateAskFriends(askFriends)
			if err != nil {
				resp = NewResponseErr(err)
				writeResponse(w, resp, http.StatusInternalServerError)
				return
			}
		case ActionRefuse:
			// notify the asker.
			askFriends.RefuseByReceiver()
			err = rout.DB.UpdateAskFriends(askFriends)
			if err != nil {
				resp = NewResponseErr(err)
				writeResponse(w, resp, http.StatusInternalServerError)
				return
			}
		case ActionIgnore:
			// TODO
		}

		resp.OK = true
		writeResponse(w, resp, http.StatusOK)
	})
}
