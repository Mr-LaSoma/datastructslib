package hashmap

import "github.com/mr-lasoma/datastructslib/src/utils"

type HashMap[K comparable, V any] struct {
	buckets    [_BUCKETS_COUNT]bucketMap[K, V]
	globalSize int
}

// Put is the function used to put a value in the
// corresponding position (see bucketMap.Put).
func (h *HashMap[K, V]) Put(key K, value V) {
	if h.buckets[computeBucketsIndex(key)].Put(key, value) {
		h.globalSize++
	}
}

// Get is the function used to get a value by the corresponding key.
// If the key is not found it returns an error (see bucketMap.Get).
func (h *HashMap[K, V]) Get(key K) (V, error) {
	return h.buckets[computeBucketsIndex(key)].Get(key)
}

// Remove is the function used to remove a key and its corrisponding value (it returns the value aswell).
// If the key is not found it returns an error (see bucketMap.Remove)
func (h *HashMap[K, V]) Remove(key K) (V, error) {
	val, err := h.buckets[computeBucketsIndex(key)].Remove(key)
	if err == nil {
		h.globalSize--
	}
	return val, err
}

// Clear is the function used to reset the HashMap (see bucketMap.Clear).
func (h *HashMap[K, V]) Clear() {
	for _, v := range h.buckets {
		v.Clear()
	}
	h.globalSize = 0
}

// ToSlice is the function used to get a slice based of all the values of the HashMap.
// This function checks every bucket, so it might be slow! (see bucketMap.ToSlice)
func (h *HashMap[K, V]) ToSlice() []V {
	slice := make([]V, h.globalSize)
	for _, v := range h.buckets {
		slice = append(slice, v.ToSlice()...)
	}
	return slice
}

// IsEmpty is the function used to check if all the buckets are empty.
func (h *HashMap[K, V]) IsEmpty() bool {
	for _, v := range h.buckets {
		if !v.IsEmpty() {
			return false
		}
	}
	return true
}

// computeBucketsIndex is a function to find the bucket where
// a key-value pair must be stored.
func computeBucketsIndex[K comparable](key K) int {
	return int(utils.MustHashCode(key)) % _BUCKET_SIZE
}
