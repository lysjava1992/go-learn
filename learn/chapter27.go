package main

import "fmt"

type IPAddr [4]byte

func (t IPAddr) String() string {

	return fmt.Sprintf("%v.%v.%v.%v", t[0], t[1], t[2], t[3])
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
