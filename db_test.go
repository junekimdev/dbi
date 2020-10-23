package dbi

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func testSetup(t *testing.T) func() {
	t.Log("Set Up")
	err := godotenv.Load(".env")
	if err != nil {
		t.Errorf("Error loading .env file: %v\n", err)
	}
	Connect(CreateURIFromEnv().String())
	return func() {
		t.Log("Tear Down")
	}
}

func TestConnect(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Errorf("Error loading .env file: %v\n", err)
	}

	Connect(CreateURIFromEnv().String())

	if db == nil {
		t.Error("Failed to connect to DB")
	}
}

// Intergration test
func TestQueryFlow(t *testing.T) {
	teardown := testSetup(t)
	defer teardown()

	expect := os.Getenv("PGUSER")

	// Query the db
	result := Query("SELECT current_user")

	var dest string

	// Scan the query result
	if err := Scan(result, func() { result.Scan(&dest) }); err != nil {
		t.Error(err)
	}

	if dest != expect {
		t.Errorf("Query result are not as expected: want-%v got-%v", expect, dest)
	}
}
