package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2}
	for index, value := range arr {
		fmt.Printf(" index: %d ;value: %d \n", index, value)
	}
	var siteMap = make(map[string]int16)
	siteMap["k1"] = 1
	k2 := "k2"
	siteMap[k2] = 2
	for key, value := range siteMap {
		fmt.Printf("key : %s ;value: %d \n", key, value)
	}
	value, ok := siteMap["k3"]
	if ok {
		fmt.Printf("value: %d \n", value)
	} else {
		fmt.Printf("key: k3  不存在\n")
	}
}
