package utils

import (
	"encoding/json"
	"hash/fnv"
)

// HashCode is a function to obtain the hashcode
// of any type. If the json.Marshal fails this methods
// returns an error
func HashCode[T any](v T) (uint64, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}

	h := fnv.New64a()
	h.Write(b)
	return h.Sum64(), nil
}

// MustHashCode is a function to obtain the hashcode of any type.
// If the json.Marshal fails this methods panics the programm
func MustHashCode[T any](v T) uint64 {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	h := fnv.New64a()
	h.Write(b)
	return h.Sum64()
}
