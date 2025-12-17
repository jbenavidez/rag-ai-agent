package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *Repository) readJSON(w http.ResponseWriter, r *http.Request, data any) error {

	maxBytes := 1024 * 1024 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields() // optional

	err := dec.Decode(data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must only contian single json value")
	}
	return nil

}

func (m *Repository) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	out, err := json.Marshal(data)

	if err != nil {
		return err
	}
	//add header if provided
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}
