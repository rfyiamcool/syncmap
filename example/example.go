package main

import (
	"fmt"

	"github.com/rfyiamcool/syncmap"
)

func main() {
	list := map[string]interface{}{
		"a": "aaa",
		"b": "bbb",
		"c": "ccc",
	}

	var m syncmap.Map
	for k, v := range list {
		m.Store(k, v)
	}

	go func() {
		m.Delete("c")
		m.Store("d", "ddd")
	}()

	for k, v := range list {
		m.Store(k, v)
	}

	fmt.Println("range println")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	fmt.Println("length: ", *m.Length())
}
