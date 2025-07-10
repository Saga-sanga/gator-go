package config

import "testing"

func TestRead(t *testing.T) {
	got, err := Read()
	if err != nil {
		t.Errorf("Failed to read file: %q", err)
	}
	want := Config{
		DBURL:           "postgres://example",
		CurrentUserName: "sanga",
	}

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func TestWrite(t *testing.T) {
	want := Config{
		DBURL:           "postgres://example",
		CurrentUserName: "sanga",
	}
	err := write(want)
	if err != nil {
		t.Errorf("Error writing: %v", err)
	}

	got, _ := Read()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
