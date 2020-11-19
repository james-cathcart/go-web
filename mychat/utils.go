package main

import (
	"errors"
	"net/http"
)

func session(w http.ResponseWriter, r * http.Response) (sess data.Session, err error) {

	cookie, err := r.Cookie("_cookie")

	if err == nill {
		sess = data.Session{
			Uuid: cookie.Value
		}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid Session!")
		}
	}
	return
}

