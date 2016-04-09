package main

import (
	"testing"

	"os"
)

// Test diff xlsx1
func TestDifferXlsx1(t *testing.T) {
	os.Args = []string{
		"differ-xlsx",
		"-vv",
		"./testdocs/test1_org.xlsx",
		"./testdocs/test1_diff.xlsx",
	}

	app := NewApp()
	if err := app.Run(os.Args); err != nil {
		t.Error(err)
	}
}

// Test diff xlsx2
func TestDifferXlsx2(t *testing.T) {
	os.Args = []string{
		"differ-xlsx",
		"-vv",
		"./testdocs/test2_org.xlsx",
		"./testdocs/test2_diff.xlsx",
	}

	app := NewApp()
	if err := app.Run(os.Args); err != nil {
		t.Error(err)
	}
}

// Test diff xlsx3
func TestDifferXlsx3(t *testing.T) {
	os.Args = []string{
		"differ-xlsx",
		"-vv",
		"./testdocs/test3_org.xlsx",
		"./testdocs/test3_diff.xlsx",
	}

	app := NewApp()
	if err := app.Run(os.Args); err != nil {
		t.Error(err)
	}
}
