package util

import (
	"encoding/base64"
	"encoding/binary"
)

const LEN_HASH = 8 // needs to be a multiple of 4
const CIPHER_SHIFT = 13

// cipher ciphers inplace a byte vector using a Caesar like cipher
func cipher(vec []byte) []byte {
	for idx := range vec {
		val := int(vec[idx]) + CIPHER_SHIFT + idx
		vec[idx] = byte(val % 256)
	}
	return vec
}

// decipher reverts inplace the Caesar like cipher of util.cipher
func decipher(vec []byte) []byte {
	for idx := range vec {
		val := int(vec[idx]) - CIPHER_SHIFT - idx
		vec[idx] = byte((val + 256) % 256)
	}
	return vec
}

// ToHash converts an ID integer to a base64-encoded string
func ToHash(id uint64) (string, error) {
	encoded := make([]byte, binary.MaxVarintLen64)
	binary.LittleEndian.PutUint64(encoded, id)
	return base64.StdEncoding.EncodeToString(cipher(encoded))[:LEN_HASH], nil
}

// ToID converts a base64-encoded string back to an ID integer
func ToID(hash string) (uint64, error) {
	// requires the input length to be a multiple of four
	decoded, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return 0, err
	}
	decoded = decipher(decoded)

	// LittleEndian requires length of binary.MaxVarintLen64
	for len(decoded) < binary.MaxVarintLen64 {
		decoded = append(decoded, 0)
	}

	return binary.LittleEndian.Uint64(decoded), nil
}

// IsValidHash validates a given string hash using bijection
func IsValidHash(hash string) bool {
	if len(hash) != LEN_HASH {
		return false
	}
	id, err := ToID(hash)
	if err != nil {
		return false
	}
	rehash, err := ToHash(id)
	if err != nil {
		return false
	}
	return hash == rehash
}
