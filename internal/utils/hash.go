package util

import (
	"encoding/base64"
	"strconv"
)

// ToHash converts an integer ID to a base64-encoded string
func ToHash(id int64) string {
	// Convert the integer ID to a string
	idStr := strconv.FormatInt(id, 10)

	// Convert the string to a byte slice
	idBytes := []byte(idStr)

	// Encode the byte slice to base64 string
	encodedID := base64.StdEncoding.EncodeToString(idBytes)

	return encodedID
}

// ToID converts a base64-encoded string back to an integer
func ToID(encodedID string) (int64, error) {
	// Decode the base64 string back to a byte slice
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedID)
	if err != nil {
		return 0, err // Return an error if decoding fails
	}

	// Convert the byte slice back to a string
	decodedStr := string(decodedBytes)

	// Convert the string back to an integer
	id, err := strconv.ParseInt(decodedStr, 10, 64)
	if err != nil {
		return 0, err // Return an error if conversion to int fails
	}

	return id, nil
}
