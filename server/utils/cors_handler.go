package utils

import "net/http"

type Handler func(http.ResponseWriter, *http.Request) *Error

type Error struct {
	Error   error
	Message string
	Code    int
}

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}
