// tests/app_test.go

package tests

import (
	"testing"
	// import path from "github.com/goobric/"
)

func TestAppInitialization(t *testing.T){
	// Initialize the application
	appInstance := &app.App{}
	appInstance.Initialize("database-url", nil)

	// Perform test assertions
	if appInstance == nil {
		t.Error("Expected app instance to be initialized, got nil!")
	}
}