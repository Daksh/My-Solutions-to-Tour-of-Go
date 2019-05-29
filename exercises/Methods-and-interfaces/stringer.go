package main

import "fmt"
//import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string{
	return fmt.Sprintf("%d.%d.%d.%d", i[0],i[1],i[2],i[3])
	//	return strconv.Itoa(int(i[0]))+"."+strconv.Itoa(int(i[1]))+"."+strconv.Itoa(int(i[2]))+"."+strconv.Itoa(int(i[3]))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
//		fmt.Printf("%T: %v\n",ip[0],int(ip[0]))
		fmt.Printf("%v: %v\n", name, ip)
	}
}
