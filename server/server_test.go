package main

import (
	"fmt"
	"testing"
)

// see https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

// function testing getLeaderboard(db *sql.DB) in db.go
func TestGetLeaderboard(t *testing.T) {

	db, err := connect_db("../data/test.sqlite") // open test database
	if err != nil {
		fmt.Errorf("could not connect to database: %w", err)
	}
	create_table(db) // create leaderboard table in that database

	// deletes the leaderboard table from the test database after the tests are run
	defer func() {
		db.Exec("DROP TABLE leaderboard")
		db.Close()
	}()

	// add test data to database
	add_user("Amoniker", db)
	add_user("kim", db)
	add_user("harry", db)
	increment_wins("Amoniker", db)

	got := getLeaderboard(db)
	// expected leaderboard values for the msgStruct
	entries := []LB_Entry{LB_Entry{Username: "Amoniker", Wins: 1}, LB_Entry{Username: "kim", Wins: 0}, LB_Entry{Username: "harry", Wins: 0}}
	want := msgStruct{MsgType: "leaderboard", Leaderboard: entries}

	// compare what getLeaderboard() got to what we want
	if got.MsgType != want.MsgType {
		t.Errorf("MsgType: got %q, wanted %q", got.MsgType, want.MsgType)
	}
	for i := 0; i < 3; i++ {
		if got.Leaderboard[i] != want.Leaderboard[i] {
			t.Errorf("Leaderboard[%d]: got %q, wanted %q", i, got.Leaderboard[i], want.Leaderboard[i])
		}
	}
}
