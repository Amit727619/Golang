// create map using make function
// Access Map Elements
// Update and Add Map Elements
// Remove Element from Map
// Iterate Over Maps

package main

import "fmt"

func main() {
    myMap := make(map[string]int)
   
    myMap["one"] = 1
    myMap["two"] = 2
    myMap["three"] = 3
    myMap["four"] = 4

    fmt.Println("value at 'two' ", myMap["two"])


    if value, exists := myMap["four"]; exists {
        fmt.Println("Value for key 'four':", value)
    } else {
        fmt.Println("Key 'four' not found")
    }

    fmt.Println("Iterat the map:")
    for key, value := range myMap {
        fmt.Printf("%s: %d\n", key, value)
    }

    delete(myMap, "two")
    fmt.Println("After deleting key 'two' ", myMap)  
	
}

