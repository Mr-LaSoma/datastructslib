package hashmap

import (
	"fmt"

	"github.com/mr-lasoma/datastructslib/src/utils"
)

// bucketMap is the sub struct for the custom hashmap implementation.
type bucketMap[K comparable, V any] struct {
	// topHashCodes are used to make the comparations (when searching from a key) faster
	topHashCodes [_BUCKET_SIZE]int

	// keys are the variables responsible for identifying each entry in the hashmap.
	// Each key is associated with a value and is stored in a bucket determined
	// by the hash function. Collisions are handled within the same bucket.
	keys [_BUCKET_SIZE]K

	// values are the values...
	values [_BUCKET_SIZE]V

	// overflow is the structure used to handle many hash collisions.
	// Each entry in bucketMap stores additional key-value pairs
	overflow *bucketMap[K, V]

	// currentSize is the current empty index of hashcodes, keys and values
	currentSize int

	// defaultKeyValue is the zero value of the K type
	defaultKeyValue K

	// defaultValValue is the zero value of the V type
	defaultValValue V
}

// newBucketMap is the base function to get a [K and V] default bucketMap
func newBucketMap[K comparable, V any]() *bucketMap[K, V] {
	return &bucketMap[K, V]{}
}

// Put is the function used to put a value in the
// corresponding key index. If there is no space, it initialize and
// puts the key in the overflow (recursively)
func (b *bucketMap[K, V]) Put(key K, value V) {
	th := computeTopHashCode(key)
	indx, err := b.findKeyIndex(key, th)
	if err == nil {
		b.values[indx] = value
		return
	}

	if b.currentSize >= len(b.keys) {
		if b.overflow == nil {
			b.overflow = newBucketMap[K, V]()
		}
		b.overflow.Put(key, value)
		return
	}

	b.topHashCodes[b.currentSize] = th
	b.keys[b.currentSize] = key
	b.values[b.currentSize] = value
	b.currentSize++
}

// Get is the function used to get a value by the corresponding key.
// If the key is not found, it searches in the overflow (recursively), until either
// the overflow is nil (returns an error) or the key is found.
func (b *bucketMap[K, V]) Get(key K) (V, error) {
	if b.IsEmpty() {
		return b.defaultValValue, fmt.Errorf("Can't get value from non existing key %#v", key)
	}

	th := computeTopHashCode(key)
	indx, err := b.findKeyIndex(key, th)
	if err == nil {
		return b.values[indx], nil
	}

	if b.overflow == nil {
		return b.defaultValValue, fmt.Errorf("Can't get value from non existing key %#v", key)
	}

	return b.overflow.Get(key)
}

// remove is the function used to remove a key and its corrisponding value.
// If the key is not found, it searches in the overflow (recursively), until either
// the overflow is nil (returns an error) or the key is found.
func (b *bucketMap[K, V]) Remove(key K) (V, error) {
	if b.IsEmpty() {
		return b.defaultValValue, fmt.Errorf("Can't remove non existing key %#v", key)
	}

	th := computeTopHashCode(key)
	indx, err := b.findKeyIndex(key, th)
	if err == nil {
		return b.removeAndShift(indx), nil
	}

	if b.overflow == nil {
		return b.defaultValValue, fmt.Errorf("Can't remove non existing key %#v", key)
	}

	return b.overflow.Remove(key)
}

// clear is the function used to reset the bucketMap.
func (b *bucketMap[K, V]) Clear() {
	b = newBucketMap[K, V]()
}

// computeTopHash is a function to compute the tophash from a
// comparable value (see bucketMap.findKeyIndex).
func computeTopHashCode[K comparable](key K) int {
	return int(utils.MustHashCode(key) >> 16)
}

// findKeyIndex is a function used to find the index of a given key
// in the bucketMap. if the key is not found then it returns an error.
func (b *bucketMap[K, V]) findKeyIndex(key K, topHash int) (int, error) {
	for i := range b.currentSize {
		if b.topHashCodes[i] == topHash && b.keys[i] == key {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Key %#v not found in bucketMap", key)
}

// removeAndShift is the helper function of bucketMap.remove, it remove and returns
// the value at the indx position and it shifts all other elements
func (b *bucketMap[K, V]) removeAndShift(indx int) V {
	value := b.values[indx]

	for i := indx + 1; i < b.currentSize; i++ {
		b.topHashCodes[i-1] = b.topHashCodes[i]
		b.keys[i-1] = b.keys[i]
		b.values[i-1] = b.values[i]
	}

	b.topHashCodes[b.currentSize-1] = 0
	b.keys[b.currentSize-1] = b.defaultKeyValue
	b.values[b.currentSize-1] = b.defaultValValue

	b.currentSize--
	return value
}

// IsEmpty is the function used to check if both the current bucketMap and all the overflows are empty.
func (b *bucketMap[K, V]) IsEmpty() bool {
	return b.currentSize == 0 && (b.overflow == nil || b.overflow.IsEmpty())
}
