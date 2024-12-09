package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, c int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(c)
	enc := json.NewEncoder(w)
	if err := enc.Encode(v); err != nil {
		log.Printf("WriteJSON_Error: %v", err)
		return err
	}
	return nil
}

func ParseJSON(r *http.Request, v interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}
