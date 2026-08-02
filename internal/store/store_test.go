package store

import "testing"

func TestSaveAndLoad(t *testing.T) {
	configDir = t.TempDir()
	defer func() { configDir = "" }()
	binaries := []Binary{
		{Alias: "java17", Path: "/usr/lib/jvm/java-17/bin/java"},
		{Alias: "java8", Path: "/usr/lib/jvm/java-8/bin/java"},
	}

	err := Save(binaries)
	if err != nil {
		t.Fatalf("Save() returned an error: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() returned an error: %v", err)
	}

	if len(loaded) != len(binaries) {
		t.Fatalf("expected %d binaries, got %d", len(binaries), len(loaded))
	}

	if loaded[0].Alias != "java17" || loaded[0].Path != "/usr/lib/jvm/java-17/bin/java" {
		t.Errorf("first binary doesn't match, got %+v", loaded[0])
	}

	if loaded[1].Alias != "java8" || loaded[1].Path != "/usr/lib/jvm/java-8/bin/java" {
		t.Errorf("second binary doesn't match, got %+v", loaded[1])
	}
}

func TestLoadWhenFileDoesNotExist(t *testing.T) {
	configDir = t.TempDir()
	defer func() { configDir = "" }()

	loaded, err := Load()
	if err != nil {
		t.Fatalf("expected no error when file doesn't exist, got: %v", err)
	}

	if len(loaded) != 0 {
		t.Errorf("expected empty slice, got %d items", len(loaded))
	}
}
