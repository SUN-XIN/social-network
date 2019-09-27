package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerGetProfiles fetches a list of profiles by the given profile IDs.
func HandlerGetProfiles(rout *Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *GetProfilesResponse

		// Fetch session and valide it.
		sessID := r.URL.Query().Get("session")

		session, err := rout.DB.GetSession(sessID)
		if err != nil {
			resp = NewGetProfilesResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if err = session.Validate(); err != nil {
			resp = NewGetProfilesResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		// Parse request from body.
		defer r.Body.Close()
		var targetProfileIDs []string
		if err = json.NewDecoder(r.Body).Decode(&targetProfileIDs); err != nil {
			resp = NewGetProfilesResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if len(targetProfileIDs) == 0 {
			resp = NewGetProfilesResponseErr(fmt.Errorf("target profile ids are required"))
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		targetProfiles, err := rout.DB.GetProfiles(targetProfileIDs)
		if err != nil {
			resp = NewGetProfilesResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if len(targetProfileIDs) != len(targetProfiles) {
			resp = NewGetProfilesResponseErr(fmt.Errorf("only %d profiles are fetched for %d ids", len(targetProfiles), len(targetProfileIDs)))
			writeResponse(w, resp, http.StatusPartialContent)
			return
		}

		// build response.
		resp = NewGetProfilesResponseOK(targetProfiles)
		writeResponse(w, resp, http.StatusOK)
	})
}
