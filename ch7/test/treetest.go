package main

import (
	"fmt"
	"gobook/ch7/newtreesort"
)

func main() {
	var t *newtreesort.Tree
	t = newtreesort.Add(t, 10)
	t = newtreesort.Add(t, 20)
	t = newtreesort.Add(t, 50)
	t = newtreesort.Add(t, 30)
	t = newtreesort.Add(t, 6)
	t = newtreesort.Add(t, 5)
	t = newtreesort.Add(t, 15)
	t = newtreesort.Add(t, 25)
	t = newtreesort.Add(t, 8)

	fmt.Println(t.ToString())

}
