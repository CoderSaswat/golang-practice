package main

import (
	"fmt"
	"sync"
)

var nm = "saswat"
var m sync.Mutex

func updateName(n string, s *sync.Mutex) {
	s.Lock()
	nm = n
	s.Unlock()
}
func main() {
	updateName("sk", &m)
	fmt.Println(nm)
	updateName("gh", &m)
	fmt.Println(nm)
}
