package hashtables

type HashTable struct {
	items map[int]int
}

/**
This function returns a simple hash by looping over the key length.
We assume this to be O(1)
*/
func hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = hash*31 + int(key[i])
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
This function returns a value for a hash
O(1)
*/
func (ht *HashTable) Get(key string) int {
	address := hash(key)
	currentBucket := ht.items[address]
	return currentBucket
}

/**
This function returns the size of the hashtable
*/
func (ht *HashTable) Size() int {
	size := len(ht.items)

	return size
}

// func main() {
// 	hashTable := &HashTable{}
// 	hashTable.Set("Grapes", 10000)
// 	hashTable.Set("Apples", 800)
// 	hashTable.Set("Oranges", 200)
// 	hashTable.Get("Grapes")
// }
