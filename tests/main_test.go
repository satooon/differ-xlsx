package tests

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
