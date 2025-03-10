package main

import (
	"database/sql"
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

	// create test data
	createTestData(db)

	got := getLeaderboard(db)
	// expected leaderboard values for the msgStruct
	entries := []LB_Entry{
		{Username: "wilford", Wins: 16},
		{Username: "jordan", Wins: 14},
		{Username: "gaylord", Wins: 12},
		{Username: "ambrose", Wins: 10},
		{Username: "freddy", Wins: 9},
		{Username: "craig", Wins: 8},
		{Username: "antonia", Wins: 6},
		{Username: "harry", Wins: 4},
		{Username: "sally", Wins: 3},
		{Username: "Amoniker", Wins: 1},
		{Username: "kim", Wins: 0},
	}

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

// fill the database with test data
func createTestData(db *sql.DB) {
	// add test data to database
	add_user("Amoniker", db)
	add_user("kim", db)
	add_user("harry", db)
	add_user("sally", db)
	add_user("ambrose", db)
	add_user("antonia", db)
	add_user("freddy", db)
	add_user("gaylord", db)
	add_user("jordan", db)
	add_user("craig", db)
	add_user("wilford", db)

	increment_wins("Amoniker", db)

	for i := 0; i < 4; i++ {
		increment_wins("harry", db)
	}

	for i := 0; i < 3; i++ {
		increment_wins("sally", db)
	}

	for i := 0; i < 10; i++ {
		increment_wins("ambrose", db)
	}

	for i := 0; i < 6; i++ {
		increment_wins("antonia", db)
	}

	for i := 0; i < 9; i++ {
		increment_wins("freddy", db)
	}

	for i := 0; i < 12; i++ {
		increment_wins("gaylord", db)
	}

	for i := 0; i < 14; i++ {
		increment_wins("jordan", db)
	}

	for i := 0; i < 8; i++ {
		increment_wins("craig", db)
	}

	for i := 0; i < 16; i++ {
		increment_wins("wilford", db)
	}
}
