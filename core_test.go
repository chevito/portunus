package main

import (
	"errors"
	"testing"
)

func TestPut(t *testing.T) {
	err := Put("animal", "dog")

	if err != nil {
		t.Errorf("Could not add value: %s", err)
	}

	value, err := Get("animal")

	if err != nil {
		t.Errorf("Key not found")
	}

	if value != "dog" {
		t.Errorf("Expected 'dog' got %s", value)
	}

}

func TestDelete(t *testing.T) {
	Put("animal", "dog")
	err := Delete("animal")
	if err != nil {
		t.Errorf("Could not delete key: %s", err)
	}

	_, err = Get("animal")
	if err == nil {
		t.Errorf("Expected error was not thrown.")
	}

	if !errors.Is(err, ErrorNoSuchKey) {
		t.Errorf("Expecting 'ErrorNoSuchKey' but received another type")
	}

}
