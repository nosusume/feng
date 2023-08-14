package feng

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SetenvMap sets environment variables based on the provided map.
//
// Takes in a map of string key-value pairs representing environment variables.
// Returns an error if there is an issue setting any of the environment variables.
func SetenvMap(envMap map[string]string) error {
	for key, value := range envMap {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}
	return nil
}

// GetenvMap retrieves a map of environment variables with a given prefix.
//
// It takes a prefix string as a parameter and returns a map[string]string
// containing the environment variables that have keys starting with the given
// prefix. If the prefix is an empty string, it retrieves all environment
// variables.
func GetenvMap(prefix string) map[string]string {
	// Get all environment variables
	envs := os.Environ()

	// Create a map to store the resulting key-value pairs
	envMap := make(map[string]string)

	// Iterate through each environment variable
	for _, v := range envs {
		// Split the variable into key-value pair
		envLine := strings.Split(v, "=")
		k := envLine[0]
		v := envLine[1]

		// Check if the key starts with the given prefix or prefix is empty
		if strings.HasPrefix(k, prefix) || prefix == "" {
			// Add the key-value pair to the map
			envMap[k] = v
		}
	}

	// Return the resulting map
	return envMap
}

// ReadEnvFile reads the contents of a .env file into a map
// Args:
//
//	file (string): The path to the .env file
//
// Returns:
//
//	map[string]string: A map containing the key-value pairs from the .env file
//	error: An error if there was a problem reading the file
func ReadEnvFile(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envMap[key] = value
		}
	}

	return envMap, nil
}

// WriteEnvFile writes the contents of a map to a .env file
//
// The function takes a prefix string and a filename string as parameters.
// It retrieves a map of environment variables using the GetenvMap function.
// If the map is empty, the function returns nil.
// Otherwise, it creates a new file with the given filename and writes each
// key-value pair from the map to the file in the format "key=value\n".
// Finally, it returns nil if the file is successfully written, or an error
// if any error occurs during the process.
func WriteEnvFile(prefix string, filename string) error {
	// Retrieve the environment variable map
	envMap := GetenvMap(prefix)

	// Return nil if the map is empty
	if len(envMap) == 0 {
		return nil
	}

	// Create a new file with the given filename
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create a buffered writer for the file
	w := bufio.NewWriter(f)

	// Write each key-value pair from the map to the file
	for k, v := range envMap {
		_, err := w.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return err
		}
	}

	// Flush the buffer and check for any error
	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

// ClearEnvSetting clears environment settings for the given environment names.
// It takes a variadic parameter of environment names and returns an error, if any.
func ClearEnvSetting(envNames ...string) error {
	// Iterate over each environment name
	for _, name := range envNames {
		// Unset the environment variable
		if err := os.Unsetenv(name); err != nil {
			return err
		}
	}
	// Return nil if no error occurred
	return nil
}
