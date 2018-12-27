package main

import (
	"fmt"
	"time"

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
		deleteOk := m.Delete("c")
		fmt.Println("delete ok ? ", deleteOk)

		deleteOk = m.Delete("d")
		fmt.Println("delete ok ? ", deleteOk)
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

	time.Sleep(1 * time.Second)
}
