package feng

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	lineRegx = regexp.MustCompile(`\A\s*(?:export\s+)?([\w\.]+)(?:\s*=\s*|:\s+?)('(?:\'|[^'])*'|"(?:\"|[^"])*"|[^#\n]+)?\s*(?:\s*\#.*)?\z`)
	// TODO: Handle variable environment variables
	// variableRegx = regexp.MustCompile(`(\\)?(\$)(\{?([A-Z0-9_]+)?\}?)`)
	// unescapeRgx  = regexp.MustCompile(`\\([^$])`)
)

// GetenvInt8 retrieves the value of the specified environment variable as an int8.
//
// It takes a string parameter `key` which specifies the name of the environment variable to retrieve.
//
// The function returns an int8 and an error. The int8 represents the value of the environment variable
// converted to int8. The error is non-nil if there was an error retrieving or converting the value.
func GetenvInt8(key string) (int8, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("environment variable not found: %s", key)
	}

	intValue, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("failed to convert environment variable to int8: %s", key)
	}

	return int8(intValue), nil
}

// GetenvInt16 retrieves the value of the specified environment variable and converts it to an int16.
//
// Parameters:
// - key: The name of the environment variable.
//
// Returns:
// - int16: The value of the environment variable as an int16.
// - error: An error if the environment variable is not set or if it fails to be parsed as an int16.
func GetenvInt16(key string) (int16, error) {
	val := os.Getenv(key)
	if val == "" {
		return 0, fmt.Errorf("environment variable %s not set", key)
	}
	num, err := strconv.ParseInt(val, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("failed to parse environment variable %s as int16: %w", key, err)
	}
	return int16(num), nil
}

// GetenvInt64 retrieves the value of the environment variable specified by the key parameter and returns it as an int64.
//
// Parameters:
// - key: The name of the environment variable.
//
// Returns:
// - int64: The value of the environment variable as an int64.
// - error: An error if the environment variable does not exist or if it cannot be parsed as an int64.
func GetenvInt64(key string) (int64, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, nil
	}
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsedValue, nil
}

// GetenvInt32 returns the integer value of the environment variable with the given key.
//
// Parameters:
// - key: the key for the environment variable.
//
// Returns:
// - int32: the integer value of the environment variable, or 0 if the variable is not set or cannot be parsed as an integer.
func GetenvInt32(key string) (int32, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return 0, nil
	}
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(intValue), nil
}

// GetenvUint8 is a function that retrieves and converts an environment variable to an unsigned 8-bit integer.
//
// It takes a string parameter `key` which represents the name of the environment variable to retrieve.
//
// It returns a uint8 value, which is the converted value of the environment variable, and an error if the conversion fails.
func GetenvUint8(key string) (uint8, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, errors.New("environment variable not set")
	}
	i, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("failed to parse environment variable: %w", err)
	}
	return uint8(i), nil
}

// GetenvUint16 retrieves the value of the environment variable named by the key
// parameter and returns it as a uint16. If the environment variable is not set
// or if the value cannot be parsed as a uint16, it returns an error.
func GetenvUint16(key string) (uint16, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, errors.New("environment variable not set")
	}

	i, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}

	return uint16(i), nil
}

// GetenvFloat64 returns the float64 value of the environment variable specified by the key parameter.
func GetenvFloat64(key string) (float64, error) {
	valueStr, ok := os.LookupEnv(key)
	if !ok {
		return 0, fmt.Errorf("%s environment variable not set", key)
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s environment variable as float64: %w", key, err)
	}
	return value, nil
}

// GetenvFloat32 retrieves the value of the environment variable with the specified key and converts it to a float32.
//
// Parameters:
// - key: The key of the environment variable to retrieve.
//
// Returns:
// - float32: The value of the environment variable, converted to a float32.
// - error: An error if the conversion fails or the environment variable does not exist.
func GetenvFloat32(key string) (float32, error) {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseFloat(valueStr, 32)
	if err != nil {
		return 0, err
	}

	return float32(value), nil
}

// GetenvUint64 retrieves the value of the environment variable with the specified key and converts it to an unsigned 64-bit integer.
//
// Parameters:
// - key: the name of the environment variable to retrieve.
//
// Returns:
// - uint64: the value of the environment variable as an unsigned 64-bit integer.
// - error: any error that occurred during the conversion or retrieval process.
func GetenvUint64(key string) (uint64, error) {
	return strconv.ParseUint(os.Getenv(key), 10, 64)
}

// GetenvUint32 returns the value of the environment variable as a uint32.
// It returns an error if the environment variable value cannot be parsed or if it is not present.
func GetenvUint32(key string) (uint32, error) {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseUint(valueStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(value), nil
}

// GetenvInt returns an integer value from the environment variable specified by the given key.
//
// Parameters:
// - key: The name of the environment variable to retrieve the integer value from.
//
// Returns:
// - int: The integer value parsed from the environment variable.
// - error: An error if the value cannot be parsed as an integer or if the environment variable does not exist.
func GetenvInt(key string) (int, error) {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse environment variable as integer: %w", err)
	}
	return value, nil
}

// GetenvBool retrieves the boolean value of the specified environment variable.
//
// It takes a single parameter, which is the key string representing the name of the environment variable.
// The function returns a boolean value and an error.
func GetenvBool(key string) (bool, error) {
	value := os.Getenv(key) // Get the value of the environment variable
	if value == "" {        // Check if the value is empty
		return false, nil
	}

	// Convert the value to a boolean
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}

	return result, nil
}

// GetEnvOrDefault returns the value of an environment variable identified by the given key.
// If the environment variable is not found or its value is empty, the function returns the defaultValue.
//
// Parameters:
// - key: the key of the environment variable to retrieve.
// - defaultValue: the default value to return if the environment variable is not found or its value is empty.
//
// Return type:
// - string: the value of the environment variable or the defaultValue.
func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

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
	data, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]string)

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		// skip empty lines and comment line
		if l == "" || l[0] == '#' {
			continue
		}
		// trim export start
		l = strings.TrimPrefix(l, "export ")
		parts := lineRegx.FindStringSubmatch(l)
		if len(parts) != 0 {
			key := removeQuotes(strings.TrimSpace(parts[1]))
			value := removeQuotes(strings.TrimSpace(parts[2]))
			envMap[key] = value
		}
	}

	return envMap, nil
}

// Load reads an environment file and sets the environment variables accordingly.
//
// It takes a variable number of filenames as parameters and returns an error if any operation fails.
func Load(filenames ...string) error {
	// Create a map to store the environment variables
	envMap := make(map[string]string)

	// Iterate over each filename provided
	for _, filename := range filenames {
		// Read the environment file and get the temporary environment map
		tempEnvMap, err := ReadEnvFile(filename)
		if err != nil {
			// Return an error if reading the environment file fails
			return fmt.Errorf("failed to read env file: %w", err)
		}
		// Merge the temporary environment map with the main environment map
		envMap = mergeMaps(envMap, tempEnvMap)
	}

	// If no filenames are provided, read the default ".env" file
	if len(filenames) == 0 {
		tempEnvMap, err := ReadEnvFile(".env")
		if err != nil {
			// Return an error if reading the environment file fails
			return fmt.Errorf("failed to read env file: %w", err)
		}
		// Merge the temporary environment map with the main environment map
		envMap = mergeMaps(envMap, tempEnvMap)
	}

	// Set the environment variables using the map
	if err := SetenvMap(envMap); err != nil {
		// Return an error if setting the environment variables fails
		return fmt.Errorf("failed to set environment variables: %w", err)
	}

	// Return nil if there are no errors
	return nil
}

// mergeMaps merges multiple maps into a single map.
//
// The function takes in one or more maps as input and combines them into a single map.
// The input maps are passed as variadic arguments of type `map[string]string`.
//
// Returns:
// - A map of type `map[string]string` that contains the merged key-value pairs from the input maps.
func mergeMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// removeQuotes removes the quotes from the beginning and end of a string.
//
// It takes a single parameter:
// - s: the string to remove the quotes from.
//
// It returns a string.
func removeQuotes(s string) string {
	if len(s) < 2 {
		return s
	}

	firstChar := s[0]
	lastChar := s[len(s)-1]

	if (firstChar == '"' && lastChar == '"') || (firstChar == '\'' && lastChar == '\'') {
		return s[1 : len(s)-1]
	}

	return s
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
