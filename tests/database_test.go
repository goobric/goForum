// tests/database_test.go

package tests

import (
	"testing"
	//import of path from "github.com/goobric/"
)

func TestDatabaseInitialization(t *testing.T){
	// Initialize the database
	db, err := app.NewDatabase("database-url")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Perform test assertions
	if db == nil {
		t.Error("Expected database instance to be initialized, got nil!")
	}
}