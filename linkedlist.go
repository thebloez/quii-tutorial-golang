package main

import (
	"container/list"
	"fmt"
)

func main() {
	testList := list.New()
	testList.PushBack("Ryan")
	testList.PushBack("Dewi")
	testList.PushBack("TEST")
	testList.PushBack("Kanaya")

	fmt.Println(testList.Len())

	for elm := testList.Front(); elm != nil; elm = elm.Next() {
		if elm.Value == "TEST" {
			testList.Remove(elm)
		}
		fmt.Println(elm.Value)

	}

	for elm := testList.Front(); elm != nil; elm = elm.Next() {
		fmt.Println(elm.Value)
	}

}
