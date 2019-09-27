package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

// GenerateID return a string uuid.
func GenerateID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// RemoveDuplicate removes the duplicate id from the given list of ids.
func RemoveDuplicate(ids []string) []string {
	var v struct{}
	res := make([]string, 0, len(ids))
	m := make(map[string]struct{}, len(ids))
	for i := range ids {
		id := ids[i]
		if _, exist := m[id]; !exist {
			res = append(res, id)
			m[id] = v
		}
	}

	return res
}

func writeResponse(w http.ResponseWriter, obj interface{}, httpCode int) {
	log.Printf("writing data into response %+v", obj)

	b, err := json.Marshal(obj)
	if err != nil {
		log.Printf("failed marshal when write data into response: %+v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(b)
}
