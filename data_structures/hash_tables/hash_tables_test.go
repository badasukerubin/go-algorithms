package hashtables

import (
	"fmt"
	"testing"
)

var size, value, expectedSize, expectedValue int

func populateHashTable(count int, start int) *HashTable {
	hashTable := &HashTable{}
	for i := start; i < (start + count); i++ {
		hashTable.Set(fmt.Sprintf("key %d", i), i)
	}
	return hashTable
}

func TestSize(t *testing.T) {
	expectedSize = 3
	hashTable := populateHashTable(expectedSize, 0)

	if size = hashTable.Size(); size != expectedSize {
		t.Errorf("count is wrong; expected %d but got %d", expectedSize, size)
	}
}

func TestSet(t *testing.T) {
	expectedSize = 3
	hashTable := populateHashTable(3, 1)
	hashTable.Set("key 1", 0)

	if size = hashTable.Size(); size != expectedSize {
		t.Errorf("count is wrong; expected %d but got %d", expectedSize, size)
	}
}

func TestGet(t *testing.T) {
	expectedValue = 0
	hashTable := populateHashTable(3, 1)
	value = hashTable.Get("key 0")

	if value != expectedValue {
		t.Errorf("value is wrong; expected %d but got %d", expectedValue, value)
	}
}
