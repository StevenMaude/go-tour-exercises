package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// Other solutions I've seen:
// * create []string, append Itoa(int(v)), then strings.Join...
// * use Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
// (The Sprintf version is probably what the tour was angling at.)
func (i IPAddr) String() string {
	var s string
	for _, v := range i {
		s += strconv.Itoa(int(v)) + "."
	}
	return s[:len(s)-1]
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
