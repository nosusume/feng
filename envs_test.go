package feng_test

import (
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
