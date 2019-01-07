package main 

import (
	"fmt"
	"sort"
)

func main() {
	salary := make(map[string]int)
	salary["brian"] = 5000
	salary["joe"] = 32000
	salary["adam"] = 18000
	salary["fred"] = 1200
	fmt.Println("salary map contents:", salary)
	fmt.Println("salary map length: ", len(salary))

	var names []string
	for name := range salary {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Sorted by Name:")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, salary[name])
	}


}