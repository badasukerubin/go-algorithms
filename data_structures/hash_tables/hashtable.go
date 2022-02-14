package main

import "fmt"

type HashTable struct {
	items map[int]int
}

/**
This function creates a simple hash by looping over the key length.
We assume this to be O(1)
*/
func hash(key string) int {
	hash := 0
	k := fmt.Sprintf("%s", key)
	for i := 0; i < len(k); i++ {
		hash = hash*31 + int(k[i])
	}

	return hash
}

/**
This function inserts a pair of hash and values
O(1) when there's no collision
*/
func (ht *HashTable) Set(key string, value int) {
	address := hash(key)
	if ht.items == nil {
		ht.items = make(map[int]int)
	}
	ht.items[address] = value
}

/**
This function gets a value for a hash
O(1)
*/
func (ht *HashTable) Get(key string) int {
	address := hash(key)
	currentBucket := ht.items[address]
	return currentBucket
}

func (ht *HashTable) Keys() {

}

func main() {
	hashTable := &HashTable{}
	hashTable.Set("Grapes", 10000)
	hashTable.Set("Apples", 800)
	hashTable.Set("Oranges", 200)
	hashTable.Get("Grapes")
}
