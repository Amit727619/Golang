LRU Cache:

//map is a key value store without any limit on number of entries 

table := make(map[string]string)

//in KV store operation term, Put operation takes key and value as input and stores mapping
table["anurag"] = "Hyd"
table["raj"] = "Pune"

//Get operation, that takes key as input and returns the corresponding value
fmt.Println(table["anurag"])


Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
Example: 
//create cache of size 2
cache := Constructor(2)
cache.Put(1, v1)
cache.Put(2, v2)
cache.Get(1)    // returns v1
cache.Put(3, v3) // removes key 2
cache.Get(2)    // returns -1 (not found)

//defining cache as very simple key value store

type Cache struct {
    table map[string]sting
}

func (c Cache) Put(key, val string) {
    c.table[key] = val
}

func (c Cache) Get(key string) string {
    return c.table[key]
}

