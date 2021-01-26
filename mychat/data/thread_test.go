package data

import (
	"testing"
)

func ThreadDeleteAll() (err error) {
	db := db()
	defer db.Close()

	query := "delete from threads"
	_, err = db.Exec(query)
	return
}

func Test_CreateThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}

	thread, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "Cannot create thread")
	}
	if thread.UserId != users[0].Id {
		t.Error("User not linked with thread")
	}
}