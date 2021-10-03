//go:generate gotemplate "github.com/WinPooh32/gotemplate/ring" RingString(string)

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Allocate strigns buffer with capacity of 2.
	var buf = MakeRingString(2)
	buf.ForcePushBack("hello")
	buf.ForcePushBack("lovely")
	buf.ForcePushBack("world")

	var sentense = make([]string, buf.Len())
	buf.CopyTo(sentense)

	// Prints "lovely world"
	fmt.Println(strings.Join(sentense, " "))
}
