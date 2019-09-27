package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerCreateReport creates a report.
func HandlerCreateReport(rout *Router) http.Handler {
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
		var request CreateReportRequest
		if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusForbidden)
			return
		}

		if err = request.Validate(); err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusBadRequest)
			return
		}

		// Fetch target profile from db.
		profiles, err := rout.DB.GetProfiles([]string{request.TargetID})
		if err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		if len(profiles) == 0 {
			resp = NewResponseErr(fmt.Errorf("target profile is not found in db"))
			writeResponse(w, resp, http.StatusNotFound)
			return
		}

		// Insert into db.
		report := NewReport(session.ProfileID, request.TargetID, request.Motive)
		err = rout.DB.InsertReport(report)
		if err != nil {
			resp = NewResponseErr(err)
			writeResponse(w, resp, http.StatusInternalServerError)
			return
		}

		// build response.
		resp.OK = true
		writeResponse(w, resp, http.StatusOK)
	})
}
