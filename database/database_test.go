package database

import "testing"

func TestOpenAndClose(t *testing.T) {
	db, err := Open("some_db", FILEDB)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("some_db: %s\n\n", db.String())

	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
}
