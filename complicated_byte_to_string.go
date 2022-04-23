// credits : https://gotour.dev
package main

import "fmt"
import "strings"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (address IPAddr) String() string {
	var str strings.Builder
	for _, b := range address {
		fmt.Fprintf(&str, "%d.", int(b))
    }
    return str.String()[:str.Len() - 1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
