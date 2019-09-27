package main

import (
	"log"
	"net/http"
)

func main() {
	rout := NewRouter()

	//////////////////////////////////
	// ONLY FOR TEST
	e := seedGetProfiles(rout)
	if e != nil {
		panic(e)
	}
	//////////////////////////////////

	http.Handle("/profiles/get", HandlerGetProfiles(rout))
	http.Handle("/friends/submit", HandlerAskFriends(rout))
	http.Handle("/friends/list_submit", HandlerListAskFriends(rout))
	http.Handle("/friends/treate_submit", HandlerTreatAskFriends(rout))
	http.Handle("/friends/list", HandlerListFriends(rout))
	http.Handle("/report/create", HandlerCreateReport(rout))

	err := http.ListenAndServe(":8080", nil)

	log.Printf("PROG FAILED: %+v", err)
}
