package main

import "fmt"

func main() {
	t := make(map[string]string, 5)
	t["key1"] = "value1"
	t["key2"] = "value2"
	for k, v := range t {
		fmt.Printf("key=[%s],value=[%s]\n", k, v)
	}
	if u,ok := t["key2"]; ok {
		fmt.Println(u)
	}
}
