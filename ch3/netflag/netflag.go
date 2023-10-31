package main

import (
	"fmt"
	"net"
)

func isUp(v net.Flags) bool     { return v&net.FlagUp == net.FlagUp }
func turnDown(v *net.Flags)     { *v &^= net.FlagUp }
func setBroadcast(v *net.Flags) { *v |= net.FlagBroadcast }
func isCast(v net.Flags) bool   { return v&(net.FlagBroadcast|net.FlagMulticast) != 0 }

func main() {
	var v net.Flags = net.FlagMulticast | net.FlagUp
	fmt.Printf("%06b %t\n", v, isUp(v)) //010001 true
	turnDown(&v)
	fmt.Printf("%06b %t\n", v, isUp(v)) //01000 false
	setBroadcast(&v)
	fmt.Printf("%06b %t\n", v, isUp(v))   //010010 false
	fmt.Printf("%06b %t\n", v, isCast(v)) //010010 true

}
