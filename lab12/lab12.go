package lab12

import "math/rand"

type HashTable struct {
	data []interface{}
	size int
}

const k = 0.6180339887499

func(ht HashTable) New() HashTable {
	rand.Seed(10)
	return HashTable{
		data: make([]interface{}, 256),
		size: 256,
	}
}

func(ht HashTable) Hash(elem interface{}) int {
	key := rand.Int() * ht.size
	return int(float64(key) * k)
}


