package httputil

import (
	"encoding/json"
	"net/http"
)

// H400 responds with a 400 code
func H400(w http.ResponseWriter) {
	w.WriteHeader(400)
}

// H405 responds with a 405 code
func H405(w http.ResponseWriter) {
	w.WriteHeader(405)
}

// H200 responds with data
func H200(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(200)
	if data == nil {
		return
	}

	b, err := json.Marshal(data)
	if err != nil {
		panic("fail to marshal data")
	}
	w.Write(b)
}
