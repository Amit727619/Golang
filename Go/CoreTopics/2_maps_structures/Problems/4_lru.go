package main

import "fmt"

// this cache is of unbounded size, now implement a cache that can have only a fixed number of elements
// Max cache size of 10 elements

//defining cache as very simple key value store
type Cache struct {
	table  map[string]string
	numEle int
	maxEle int
}

func (c *Cache) Put(key, val string) {

	//update numEle if key is new
	_, exists := c.table[key]
	if !exists {
		if c.numEle >= c.maxEle {
			//cache has reached max capacity, can't add more values
			return
		}
		c.numEle++

	}
	fmt.Println("Put numEle:", c.numEle)

	fmt.Printf("Put myCache %p table address %p table value %p numEle address %p \n", c, &c.table, c.table, &c.numEle)

	// }
	c.table[key] = val
}

func (c *Cache) Get(key string) string {

	fmt.Println("Get numEle:", c.numEle)
	fmt.Printf("Get myCache %p table address %p table value %p numEle address %p \n", c, &c.table, c.table, &c.numEle)

	return c.table[key]
}

func main() {
	var myCache *Cache = new(Cache)
	myCache.table = make(map[string]string)
	myCache.maxEle = 10

	fmt.Printf("myCache %p table address %p table %p numEle address %p \n", myCache, &myCache.table, myCache.table, &myCache.numEle)

	myCache.Put("anurag", "HYD")
	fmt.Println("Num Ele: ", myCache.numEle)

	fmt.Println(myCache.Get("anurag"))

	myCache.Put("anurag", "Pune")

	fmt.Println("Num Ele: ", myCache.numEle)

}
