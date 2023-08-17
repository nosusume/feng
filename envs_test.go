package feng_test

import (
	"os"
	"testing"

	"github.com/nosusume/feng"
)

// compareMap compares two maps and returns true if they are equal, false otherwise.
// The maps are considered equal if they have the same keys and values.
func compareMap(m, m1 map[string]string) bool {
	// Check if the lengths of the maps are different
	if len(m) != len(m1) {
		return false
	}

	// Iterate over the keys and values of the first map
	for k, v := range m {
		// Check if the second map does not contain the key or if the value is different
		if v1, ok := m1[k]; !ok || v1 != v {
			return false
		}
	}

	// The maps are equal
	return true
}

// getMapKeys returns a slice of keys from a given map.
// The function takes in a map[string]string and returns a []string.
// It iterates over the map and appends each key to the slice.
// The resulting slice is then returned.
func getMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func TestReadEnvFile(t *testing.T) {
	// Test case 1: Reading a valid .env file
	t.Run("Reading a valid .env file", func(t *testing.T) {
		expected := map[string]string{
			"KEY1": "VALUE1",
			"KEY2": "VALUE2",
		}
		filename := ".env.test"
		// Create a test .env file with the desired content
		file, err := os.Create(filename)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer file.Close()
		_, err = file.WriteString("KEY1=\"VALUE1\"\nKEY2=\"VALUE2\"\n")
		if err != nil {
			t.Fatalf("Failed to write to test file: %v", err)
		}

		got, err := feng.ReadEnvFile(filename)
		if err != nil {
			t.Fatalf("ReadEnvFile returned an error: %v", err)
		}

		// Compare the actual result with the expected result
		for key, value := range expected {
			if got[key] != value {
				t.Errorf("Expected %s=%s, but got %s=%s", key, value, key, got[key])
			}
		}
	})

	// Test case 2: Reading an empty .env file
	t.Run("Reading an empty .env file", func(t *testing.T) {
		expected := map[string]string{}
		filename := ".env.empty"
		// Create an empty test .env file
		file, err := os.Create(filename)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer file.Close()

		got, err := feng.ReadEnvFile(filename)
		if err != nil {
			t.Fatalf("ReadEnvFile returned an error: %v", err)
		}

		// Compare the actual result with the expected result
		for key, value := range expected {
			if got[key] != value {
				t.Errorf("Expected %s=%s, but got %s=%s", key, value, key, got[key])
			}
		}
	})

	// Test case 3: Reading a .env file with comment lines
	t.Run("Reading a .env file with comment lines", func(t *testing.T) {
		expected := map[string]string{
			"KEY1": "VALUE1",
		}
		filename := ".env.comment"
		// Create a test .env file with comment lines
		file, err := os.Create(filename)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer file.Close()
		_, err = file.WriteString("# This is a comment\nKEY1=VALUE1\n# Another comment\n")
		if err != nil {
			t.Fatalf("Failed to write to test file: %v", err)
		}

		got, err := feng.ReadEnvFile(filename)
		if err != nil {
			t.Fatalf("ReadEnvFile returned an error: %v", err)
		}

		// Compare the actual result with the expected result
		for key, value := range expected {
			if got[key] != value {
				t.Errorf("Expected %s=%s, but got %s=%s", key, value, key, got[key])
			}
		}
	})
}

func TestSetenvMap(t *testing.T) {
	// Test case 1: Setting a single environment variable
	envMap1 := map[string]string{
		"KEY": "VALUE",
	}
	err := feng.SetenvMap(envMap1)
	if err != nil {
		t.Errorf("Error setting environment variables: %v", err)
	}
	m := feng.GetenvMap("KEY")
	if !compareMap(envMap1, m) {
		t.Errorf("Error setting environment variables setting result %v is different: %v", envMap1, m)
	} else {
		t.Logf("Successfully set environment variables: %v", m)

	}
	err = feng.ClearEnvSetting(getMapKeys(envMap1)...)
	if err != nil {
		t.Errorf("Error clearing environment variables: %v", err)
	}

	// Test case 2: Setting multiple environment variables
	envMap2 := map[string]string{
		"KEY1": "VALUE1",
		"KEY2": "VALUE2",
		"KEY3": "VALUE3",
	}
	err = feng.SetenvMap(envMap2)
	if err != nil {
		t.Errorf("Error setting environment variables: %v", err)
	}
	m = feng.GetenvMap("KEY")
	if !compareMap(envMap2, m) {
		t.Errorf("Error setting environment variables setting result %v is different: %v", envMap2, m)
	} else {
		t.Logf("Successfully set environment variables: %v", m)

	}
	err = feng.ClearEnvSetting(getMapKeys(envMap2)...)
	if err != nil {
		t.Errorf("Error clearing environment variables: %v", err)
	}

	// Test case 3: Setting an empty map
	envMap3 := map[string]string{}
	err = feng.SetenvMap(envMap3)
	if err != nil {
		t.Errorf("Error setting environment variables: %v", err)
	}
}
