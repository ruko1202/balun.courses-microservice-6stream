package test

import (
	"os"
	"testing"
)

var services = []string{
	"api-gateway",
	"auth",
	"chat-message",
	"friend",
	"media",
	"notification",
	"user-profile",
}

type serviceTestCases struct {
	name      string
	urlFormat string
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
