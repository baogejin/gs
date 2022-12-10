package main

import (
	"fmt"
	"gs/lib/mychooser"
)

func main() {
	c := &mychooser.MyChooser{}
	c.Add(1, 1)
	c.Add(2, 1)
	c.Add(3, 1)
	for i := 1; i < 10; i++ {
		fmt.Println(c.Pick())
	}
}
