package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseJsonFile(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		content := `
			{
				"definitions": {
					"AltLeft": {
						"timing": [
							[
								0,
								157.6
							]
						]
					}
				}
			}
		`

		path := filepath.Join(t.TempDir(), "config.json")

		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}

		config, err := ParseJsonFile(path)
		if err != nil {
			t.Fatal(err)
		}

		definition, ok := config.Definitions["AltLeft"]
		if !ok {
			t.Fatal("expected AltLeft definition")
		}

		if got := definition.Timing[0][1]; got != 157.6 {
			t.Errorf("expected 157.6, got %v", got)
		}
	})

	t.Run("file doesn't exist", func(t *testing.T) {
		_, err := ParseJsonFile(filepath.Join(t.TempDir(), "missing.json"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "config.json")

		if err := os.WriteFile(path, []byte(`{"definitions": invalid}`), 0644); err != nil {
			t.Fatal(err)
		}

		_, err := ParseJsonFile(path)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
